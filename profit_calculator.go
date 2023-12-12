package main

import (
	"fmt"
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

func main() {
	revenue := getInputFromUser("Revenue: ")
	expenses := getInputFromUser("Expenses: ")
	taxRate := getInputFromUser("Tax rate: ")

	earningsBeforeTaxes, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println("EBT: ", earningsBeforeTaxes)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)
}
