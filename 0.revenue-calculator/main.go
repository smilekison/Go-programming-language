package main

import (
	"errors"
	"fmt"
	"os"
)

func Calculate() {
	revenue, error := getUserInput("Revenue: ")
	if error != nil {
		fmt.Println(error)
		return
	}
	expenses, error := getUserInput("expenses: ")
	if error != nil {
		fmt.Println(error)
		return
	}
	tax_rate, error := getUserInput("tax_rate: ")
	if error != nil {
		fmt.Println(error)
		return
	}
	rev, exp, tax_rate := calculateFinancials(revenue, expenses, tax_rate)
	fmt.Println("This is your EBT: ", rev)
	fmt.Println("This is profit : ", exp)
	fmt.Println("This is your ratio: ", tax_rate)

	storeResults(rev, exp, tax_rate)

}

func calculateFinancials(revenue, expenses, tax_rate float64) (float64, float64, float64) {
	EBT := revenue - expenses
	EAT := EBT * (1 - tax_rate/100)
	ratio := EBT / EAT
	return EBT, EAT, ratio

}
func storeResults(rev, exp, tax_rate float64) {
	results := fmt.Sprintf("rev: %.1f\nexp: %.1f\ntax_rate:%.1f", rev, exp, tax_rate)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be positive number.")
	}
	return userInput, nil
}

func main() {
	Calculate()
}
