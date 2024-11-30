package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	outputText("Investment Calculator\n")

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future value is %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future real value is %.1f\n", futureRealValue)

	// fmt.Printf("Future value is %.3f\n", futureValue)
	// fmt.Println("Future real value is", futureRealValue)
	fmt.Printf(`Future value is %.1f\n
	Future real value is %.1f\n`, futureValue, futureRealValue)
	fmt.Println(formattedFV)
	fmt.Println(formattedRFV)

}

func outputText(text string) {
	fmt.Print(text)
}
