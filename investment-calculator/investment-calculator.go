package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 1000
	var investmentAmount float64 = 200
	expectedReturnValue := 2.2
	var years int

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter expected return value: ")
	fmt.Scan(&expectedReturnValue)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnValue/100, float64(years))
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, float64(years))
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
