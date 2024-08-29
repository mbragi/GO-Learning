package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount = 200
	var expectedReturnValue = 2.2
	var years = 5

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnValue/100, float64(years))
	fmt.Println(futureValue)
}
