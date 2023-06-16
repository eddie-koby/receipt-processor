package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eddie-koby/receipt-processor/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// unit test to confirm the process receipt endpoint returns status 200 for a valid input
func TestProcessReceipt(t *testing.T) {
	r := SetUpRouter()
	r.POST("/receipts/process", processReceipt)
	item1 := models.Item{ShortDescription: "Pepsi - 12-oz", Price: "1.25"}
	item2 := models.Item{ShortDescription: "Dasani", Price: "1.40"}
	var testReceipt models.Receipt
	testReceipt.Retailer = "Walgreens"
	testReceipt.PurchaseDate = "2022-01-02"
	testReceipt.PurchaseTime = "08:13"
	testReceipt.Total = "2.65"
	testReceipt.Items = append(testReceipt.Items, item1, item2)
	jsonValue, _ := json.Marshal(testReceipt)
	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
