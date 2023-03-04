package main

import (
	"Swap/palindrome"
	"fmt"
)

func main() {
	type Palid struct {
		a string
		b bool
	}

	var a string

	fmt.Print("so'zni kiriting: ")
	fmt.Scan(&a)
	if palindrome.Palindrome(a) {
		s := Palid{a, palindrome.Palindrome(a)}
		fmt.Println(s)
	} else {
		fmt.Println("Bu so'z palidrome emas")
	}

}
