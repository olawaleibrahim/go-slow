package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future value: %v \nFuture real valuue: %v", futureValue, futureRealValue)
	fmt.Println(futureRealValue)
	fmt.Println("Divisor result:", divisor(90, 34))
}

func divisor(denominator float64, numerator float64) float64 {
	if denominator == 0 {
		return 0
	}

	return numerator / denominator
}

func calculateFutureValues(
	investmentAmount float64, expectedReturnRate float64, years float64,
) (float64, float64) {

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, futureRealValue
}
