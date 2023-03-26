package main

import "github.com/Dostonlv/GolangBootcamp/pointers/methods"

func main() {
	var (
		s = &methods.MyStruct{
			K: &methods.K{},
			T: methods.GetPointers("Hello"),
		}
	)
	methods.PrintPointers(s)
}
