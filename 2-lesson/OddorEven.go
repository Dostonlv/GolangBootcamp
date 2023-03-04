package main

import "fmt"

func OddorEven(a int) {
	if a == 0 {
		fmt.Println("this number is 0")
	} else if a%2 == 0 {
		fmt.Println("this number is even")
	} else if a%2 != 0 {
		fmt.Println("this number is odd")
	}
}
