package swap

func Swap(a, b int) {
	//var temp int
	//temp = a
	//a = b
	//b = temp
	a, b = b, a
	println("a=", a, "b=", b)
}
