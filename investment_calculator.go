package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 2.5
	var investmentAmount float64 = 1000
	years := 10.0
	expectedReturnRate := 5.5

	//fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years) // 1000 * 1.055^10 / 1.025^10

	fmt.Printf("Future value: %f\n", futureValue)
	fmt.Printf("Future Real Value: %f\n", futureRealValue)

}
