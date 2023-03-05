package FizzBuzz

import "fmt"

func FizzBuzz(a int) {
	if a%15 == 0 {
		fmt.Println("FizzBuzz")
	} else if a%5 == 0 {
		fmt.Println("Fizz")
	} else if a%3 == 0 {
		fmt.Println("Buzz")
	}
}
