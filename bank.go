package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalaceFile = "balance.txt"

func getBalancefromfile() (float64, error) {

	data, err := os.ReadFile(accountBalaceFile)
	if err != nil {
		return 404, errors.New("Failed to read the files")

	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 405, errors.New("Failed to convert datatype")
	}
	return balance, nil

}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)

	os.WriteFile("balance.txt", []byte(balanceText), 0664)

}

func main() {
	fmt.Println("Welcome to Go Bank")
	var accountBalance, err = getBalancefromfile()
	if err != nil {
		fmt.Println("Errors")
		fmt.Println(err)
		fmt.Println("-------")
		//panic("can't continue")
	}
	fmt.Println("Welcome to Go Bank")
	for {

		communication()
		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Print("Your Balance: ", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var depositAmount float64

			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Depost Amount must be greater than zero")
				//return
				continue
			}

			accountBalance = accountBalance + depositAmount
			fmt.Println("Balance updated! NEW BALANCE : ", accountBalance)
			writeBalanceToFile(accountBalance)

		} else if choice == 3 {
			fmt.Print("Withdraw Money: ")
			var withrowAmount float64

			fmt.Scan(&withrowAmount)

			if withrowAmount <= 0 {
				fmt.Println("withrowAmount must be greater than zero")
				return
			}
			if withrowAmount > accountBalance {
				fmt.Println("withrowAmount should not be greater than accountBalance")
				return
			}
			accountBalance -= withrowAmount
			fmt.Println("Balanced Updated! New Balance : ", accountBalance)
			writeBalanceToFile(accountBalance)

		} else {
			fmt.Println("Good Bye")
			break

		}
	}
	fmt.Println("Thank you for choosing our bank")
}
