package Pyramids

func InvertedHalfPyramid(rows int) {
	for i := rows; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			print("* ")
		}
		println()
	}
}
