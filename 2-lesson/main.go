package main

import (
	"fmt"
	"lesson2/Factorial"
)

func main() {
	print("Sonni kiriting: ")
	var a int
	fmt.Scan(&a)
	//sum := 0
	//for i := 0; i <= a; i++ {
	//	if i%2 == 0 {
	//		sum += i
	//	}
	//}
	//fmt.Println(sum)
	//FizzBuzz.FizzBuzz(a)
	f := Factorial.Factorial(a)
	fmt.Println(f)
}
