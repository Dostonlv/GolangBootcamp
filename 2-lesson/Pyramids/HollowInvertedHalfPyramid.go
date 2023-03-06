package Pyramids

func HollowInvertedHalfPyramid(rows int) {
	for i := 1; i <= rows; i++ {
		for j := i; j <= rows; j++ {
			if i == 1 || j == i || j == rows {
				print("*")
			} else {
				print(" ")
			}
		}
		println()
	}
}
