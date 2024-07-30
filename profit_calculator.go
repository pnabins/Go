package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	//var revenue float64
	//var taxRate float64
	//var expenses float64
	revenue, err := getUserInput("Enter Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Enter Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Enter Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter Tax Rate: ")
	fmt.Scan(&taxRate)

	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)

	earningBeforeTax, earningAfterTAx, ratio := CalculateFinanvials(revenue, expenses, taxRate)

	//earningBeforeTax := revenue - expenses
	//earningAfterTAx := (revenue - (revenue * taxRate)) - expenses
	fmt.Printf("Earning Before Tax: %f\n", earningBeforeTax)
	fmt.Printf("Earning After Tax: %f\n", earningAfterTAx)

	//ratio := earningBeforeTax / earningAfterTAx
	fmt.Printf("Ration: %f\n", ratio)
	Storeresult(earningBeforeTax, earningAfterTAx, ratio)

}
func Storeresult(earningBeforeTax, earningAfterTAx, ratio float64) {
	results := fmt.Sprintf("earningBeforeTax: %.1f\narningAfterTAx: %.1f\nratio: %.3f\n", earningBeforeTax, earningAfterTAx, ratio)
	os.WriteFile("result.txt", []byte(results), 0644)
}

func CalculateFinanvials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningBeforeTax := revenue - expenses
	earningAfterTAx := (revenue - (revenue * taxRate)) - expenses
	ratio := earningBeforeTax / earningAfterTAx
	return earningAfterTAx, earningBeforeTax, ratio
}
func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Value should be positivie or greater than zero")
	}
	return userInput, nil

}
