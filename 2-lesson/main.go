package main

import "fmt"

func main() {
	print("Sonni kiriting: ")
	var a int
	fmt.Scan(&a)
	sum := 0
	for i := 0; i <= a; i++ {
		sum += i
	}
	fmt.Println(sum)
}
