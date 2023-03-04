package main

import (
	"Calculator/calc"
	"fmt"
)

func main() {
	var number1, number2 float64
	var operator string
	fmt.Println("Kalkulyator")
	fmt.Println("==============")
	fmt.Print("1-sonni kiriting: ")
	fmt.Scanln(&number1)
	fmt.Print("2-sonni kiriting: ")
	fmt.Scanln(&number2)
	calc.Calculator(number1, number2, operator)
}
