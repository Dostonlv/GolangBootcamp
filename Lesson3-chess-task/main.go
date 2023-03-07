package main

import (
	"chess/Chess"
	"fmt"
)

func main() {

	var x1, y1, x2, y2 int
	print("x1= ")
	fmt.Scanf("%v", &x1)
	print("y1= ")
	fmt.Scanf("%v", &y1)
	print("x2= ")
	fmt.Scanf("%v", &x2)
	print("y2= ")
	fmt.Scanf("%v", &y2)
	Chess.Chess(x1, y1, x2, y2)
}
