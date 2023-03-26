package methods

import "fmt"

func GetPointers(s string) *string {
	return &s
}

func PrintPointers(s *MyStruct) {
	fmt.Println(s)
}
