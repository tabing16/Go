package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var tax float64

	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&tax)

	ebt := revenue - expenses
	profit := ebt * (1 - tax/100)
	ratio := ebt / profit

	fmt.Println("The profit is: ", profit)
	fmt.Println("The ratio of EBT to profit is: ", ratio)
}
