package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTaxes := revenue - expenses
	profit := earningsBeforeTaxes - (earningsBeforeTaxes * taxRate)
	ratio := profit / revenue

	fmt.Println("EBT: ", earningsBeforeTaxes)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)
}
