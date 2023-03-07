package main

import (
	"bufio"
	"fmt"
	"lesson3-task/MultiplicationTable"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")
	text, _ := reader.ReadString('\n')
	var num int
	_, err := fmt.Sscan(text, &num)
	if err != nil {
		MultiplicationTable.Table()
	} else {
		MultiplicationTable.MultiplicationTable(num)
	}
}
