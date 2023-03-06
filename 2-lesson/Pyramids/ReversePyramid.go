package Pyramids

import "fmt"

func ReversePyramid(rows int) {
	for i := rows; i >= 1; i-- {
		for space := 0; space < rows-i; space++ {
			fmt.Print("  ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}
