package Fibonacci

func FibonacciSum(index int) int {
	if index <= 1 {
		return index
	}
	var sum int = 0
	var prev, curr int = 0, 1
	for i := 2; i <= index; i++ {
		next := prev + curr
		prev = curr
		curr = next
		sum += curr
	}
	return sum + 1
}
