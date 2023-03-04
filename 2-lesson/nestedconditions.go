package main

import "fmt"

func Divided_into_three(a int) {
	if a%2 != 0 {
		if a%3 == 0 {
			fmt.Println("divided into three")
		}
	}
}
