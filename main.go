package main

import (
	"github.com/eddie-koby/receipt-processor/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"net/http"
)

// Create a hashmap to store receipt ID's with their scores in memory
var scoredReceipts = make(map[string]int)

func main() {
	router := gin.Default()
	// process receipt endpoint
	router.POST("/receipts/process", processReceipt)
	// get points endpoint
	router.GET("/receipts/:id/points", getPoints)

	router.Run("localhost:8080")
}

// processReceipt endpoint
func processReceipt(c *gin.Context) {
	var newReceipt models.Receipt
	//
	if err := c.BindJSON(&newReceipt); err != nil {
		return
	}
	// get score for receipt
	score := newReceipt.CalcPoints()
	// generate unique ID for receipt
	id := uuid.New().String()
	// store ID and score pair in the hashmap
	scoredReceipts[id] = score
	// return ID
	c.IndentedJSON(http.StatusOK, gin.H{"id": id})
}

// getPoints endpoint
func getPoints(c *gin.Context) {
	// get ID from request
	id := c.Param("id")
	// check the hashmap for the given ID
	score, exists := scoredReceipts[id]
	if exists {
		c.IndentedJSON(http.StatusOK, gin.H{"points": score})
		return
	}
	// if we didn't find a matching ID then return a 404 error
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Receipt ID not found."})
}
