package main

import (
	"fmt"
	"os"
)

func getInputFromUser(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	earningsBeforeTaxes := revenue - expenses
	profit := earningsBeforeTaxes - (earningsBeforeTaxes * taxRate)
	ratio := profit / revenue

	return earningsBeforeTaxes, profit, ratio
}

func storeResults(earningsBeforeTaxes float64, profit float64, ratio float64) {
	results := fmt.Sprint("Earnings before taxes: ", earningsBeforeTaxes, "\nProfit: ", profit, "\nRatio: ", ratio)
	os.WriteFile("financials.txt", []byte(results), os.FileMode(0644))
}

func main() {
	revenue := getInputFromUser("Revenue: ")
	expenses := getInputFromUser("Expenses: ")
	taxRate := getInputFromUser("Tax rate: ")

	if revenue <= 0 || expenses <= 0 || taxRate <= 0 {
		fmt.Println("Invalid input, all values must be greater than 0.")
		return
	}

	earningsBeforeTaxes, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println("EBT: ", earningsBeforeTaxes)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)
	storeResults(earningsBeforeTaxes, profit, ratio)
}
