package MultiplicationTable

func MultiplicationTable(k int) {
	if k <= 15 {
		for i := 1; i <= 10; i++ {
			println(k, "*", i, "=", k*i)
		}
	}
}
