package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	fmt.Scan(&a)
	b := strconv.Itoa(a)
	x := 0
	for i := 0; i < len(b)-1; i++ {
		//fmt.Println(string(b[i]))

		if string(b[i]) == "1" && string(b[i+1]) == "3" {
			x = x + 1
			break
		}
	}
	if x == 0 {
		fmt.Println("omadli chipta")
	} else {
		fmt.Println("omadsiz chipta")
	}
}
