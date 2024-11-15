package comms

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

const inflationRate = 2.5
const divisorFile = "divisor.txt"

func divisor(denominator float64, numerator float64) float64 {
	if denominator == 0 {
		return 0
	} else if denominator > 1 {
		if denominator > numerator {
			return denominator / numerator
		} else {
			return numerator / denominator
		}
	} else {
		return numerator / denominator
	}

}

func invest() {

	var numerator float64
	var denominator float64
	for {
		divisorText, _ := getDivisorFromFile()
		println("Divisor from text:", divisorText)
		fmt.Print("Enter numerator: ")
		fmt.Scan(&numerator)
		fmt.Print("Enter denominator: ")
		fmt.Scan(&denominator)

		if numerator == 0 {
			break
		}
		result := divisor(denominator, numerator)
		fmt.Println("Divisor result:", result)
		writeDivisorToFile(result)
	}
}

func writeDivisorToFile(divisor float64) {
	divisorText := fmt.Sprintln(divisor)
	os.WriteFile("divisor.txt", []byte(divisorText), 0644)
}

func getDivisorFromFile() (float64, error) {
	data, err := os.ReadFile((divisorFile))

	if err != nil {
		return 1000, errors.New("failed to find divisor file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored divisor value")
	}

	return balance, nil
}

func calculateFutureValues(
	investmentAmount float64, expectedReturnRate float64, years float64,
) (float64, float64) {

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, futureRealValue
}
