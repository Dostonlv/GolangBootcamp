package Pyramids

import "fmt"

func FullPyramid(rows int) {
	for i := 1; i <= rows; i++ {
		k := 0
		for space := 1; space <= rows-i; space++ {
			fmt.Print("  ")
		}
		for {
			fmt.Print("* ")
			k++
			if k == 2*i-1 {
				break
			}
		}
		fmt.Println("")
	}
}
