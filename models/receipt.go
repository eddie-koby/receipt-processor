package models

import (
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// data structure to store info about receipts
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
}

// data structure to store info about indiviudal items
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// Function to calculate points for a receipt
func (newReceipt *Receipt) CalcPoints() int {
	score := 0

	// 1 point for each alphanumeric character
	for _, char := range newReceipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			score += 1
		}
	}

	// 5 points for every 2 items
	score += (len(newReceipt.Items) / 2) * 5

	// if item description is a multiple of 3 then award 20% (rounded up to an integer)
	// of the item price in points
	for _, item := range newReceipt.Items {
		description := strings.Trim(item.ShortDescription, " ")
		if len(description)%3 == 0 {
			itemPrice, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				panic(err)
			}
			score += int(math.Ceil(itemPrice * .2))
		}
	}

	// 50 points if the total price is a whole number
	totalPrice, err := strconv.ParseFloat(newReceipt.Total, 64)
	if err != nil {
		panic(err)
	}
	if totalPrice == math.Trunc(totalPrice) {
		score += 50
	}

	// 25 points if the total price is a multiple of 0.25
	divTotalPrice := totalPrice / 0.25
	if divTotalPrice == math.Trunc(divTotalPrice) {
		score += 25
	}

	// 6 points if the purchase date is even
	myDate, err := time.Parse("2006-01-02", newReceipt.PurchaseDate)
	if err != nil {
		panic(err)
	}
	if myDate.Day()%2 != 0 {
		score += 6
	}

	// 10 points if the purchase time is between 2 and 4 PM
	myTime, err := time.Parse("15:04", newReceipt.PurchaseTime)
	if err != nil {
		panic(err)
	}

	twoPM, _ := time.Parse("15:04:05", "14:00:00")
	fourPM, _ := time.Parse("15:04:05", "18:00:00")
	if myTime.After(twoPM) && myTime.Before(fourPM) {
		score += 10
	}

	return score
}
