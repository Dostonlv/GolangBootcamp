package calc

import "fmt"

func Calculator(number1, number2 float64, operator string) {
	fmt.Println("Amalni kiriting (+,-,*,/): ")
	fmt.Scanln(&operator)

	result := 0.0

	switch operator {
	case "+":
		result = number1 + number2
	case "-":
		result = number1 - number2
	case "*":
		result = number1 * number2
	case "/":
		result = number1 / number2
	default:
		fmt.Println("Noto'g'ri amal kiritildi")
		return
	}
	if result == float64(int(result)) {
		fmt.Printf("%d\n", int(result))
	} else {
		fmt.Printf("%.2f\n", result)
	}
}
