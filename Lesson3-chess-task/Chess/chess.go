package Chess

import (
	"fmt"
	"math"
)

func Chess(x1, y1, x2, y2 int) {
	var yDiff = math.Abs(float64(y2 - y1))
	var xDiff = math.Abs(float64(x2 - x1))

	if int(xDiff) == 1 && int(yDiff) == 2 {
		fmt.Println("Yura oladi")
	} else if int(xDiff) == 2 && int(yDiff) == 1 {
		fmt.Println("Yura oladi")
	} else {
		fmt.Println("Yura olmaydi")
	}

}
