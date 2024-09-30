// package main

// import "fmt"

// func Calculate() {
// 	var total_earning float64
// 	var tax_free float64
// 	var ni float64
// 	var paye float64

// 	fmt.Print("Enter total earning: ")
// 	fmt.Scan(&total_earning)

// 	fmt.Print("Enter tax free money: ")
// 	fmt.Scan(&tax_free)

// 	fmt.Print("Enter tax rate: ")
// 	fmt.Scan(&paye)

// 	fmt.Print("Enter national insurance rate: ")
// 	fmt.Scan(&ni)

// 	// fmt.Println("This is your total earning: ", total_earning)
// 	// fmt.Println("This is your tax free earning: ", tax_free)

// 	var taxable_income = total_earning - tax_free
// 	fmt.Println("This is your taxable_income: ", taxable_income)

// 	var total_paye_to_be_paid float64 = paye / 100 * taxable_income
// 	fmt.Println("This is total_paye_to_be_paid : ", total_paye_to_be_paid)

// 	var total_ni_to_be_paid float64 = ni / 100 * taxable_income
// 	fmt.Println("This is total_ni_to_be_paid : ", total_ni_to_be_paid)

// 	// var ration = (EBT / EAT)
// 	// fmt.Println("This is your ratio: ", ration)

// }

// func main() {
// 	Calculate()
// }
