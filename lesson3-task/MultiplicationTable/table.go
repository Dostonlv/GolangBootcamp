package MultiplicationTable

import "fmt"

func Table() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, i*j)
		}
		println()
	}
}
