package main

import "fmt"

func Condition() {
	var a, b, c int
	print("a ni kiriting: ")
	fmt.Scan(&a)
	print("b ni kiriting: ")
	fmt.Scan(&b)
	print("c ni kiriting: ")
	fmt.Scan(&c)

	if a > b && a > c {
		fmt.Printf("a = %v eng katta", a)
	} else if b > a && b > c {
		fmt.Printf("b = %v eng katta", b)
	} else if c > a && b < c {
		fmt.Printf("c = %v eng katta", c)
	} else {
		fmt.Printf("%v == %v == %v Hammasi teng", a, b, c)
	}
}
