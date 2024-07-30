package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	//
	//fmt.Print("Enter Investment Amount: ")
	outputText("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter No. of Years: ")
	fmt.Scan(&years)
	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years) // 1000 * 1.055^10 / 1.025^10

	fmt.Printf("Future value: %f\n", futureValue)
	fmt.Printf("Future Real Value: %f\n", futureRealValue)

}

func outputText(text string) {
	fmt.Print(text)
}
func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rtv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rtv = fv / math.Pow(1+inflationRate/100, years) // 1000 * 1.055^10 / 1.025^10
	return fv, rtv
}
