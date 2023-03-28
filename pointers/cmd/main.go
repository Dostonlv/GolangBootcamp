package main

import "pointers/methods"

// return address
func main() {
	var (
		s = &methods.MyStruct{
			K: &methods.K{},
			T: methods.GetPointers("Hello"),
		}
	)
	methods.PrintPointers(s)
}
