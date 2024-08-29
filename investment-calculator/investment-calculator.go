package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 200
	var expectedReturnValue float64 = 2.2
	var years int = 5

	var futureValue = investmentAmount * math.Pow(1+expectedReturnValue/100, float64(years))
	fmt.Println(futureValue)
}
