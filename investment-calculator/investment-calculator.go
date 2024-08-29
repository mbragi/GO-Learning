package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 1000
	var investmentAmount float64 = 200
	expectedReturnValue := 2.2
	years := 5

	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow(1+expectedReturnValue/100, float64(years))
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, float64(years))
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
