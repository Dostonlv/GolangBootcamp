package Pyramids

func HalfPyramid(rows int) {
	var i, j int
	for i = 1; i <= rows; i++ {
		for j = 1; j <= i; j++ {
			print("* ")
		}
		println()
	}

}
