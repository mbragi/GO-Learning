// create a profit calculator

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Profit Calculator!")
	fmt.Println("Please enter the following information:")
	//revenue
	fmt.Print("Enter revenue: ")
	var revenue float64
	fmt.Scan(&revenue)
	// expenses
	fmt.Print("Enter expenses: ")
	var expenses float64
	fmt.Scan(&expenses)
	// taxes
	fmt.Print("Enter tax Rate: ")
	var taxRate float64
	fmt.Scan(&taxRate)

	// calculate earnings before tax
	earningsBeforeTax := revenue - expenses
	// calculate earnings after tax
	profit := earningsBeforeTax * (1 - taxRate/100)
	// calculate ratio of profit and ebt
	profitRatio := profit / earningsBeforeTax
	fmt.Println("Profit: ", profit)
	fmt.Println("Earnings Before Tax: ", earningsBeforeTax)
	fmt.Println("Profit Ratio: ", profitRatio)
}
