package big_o

func factorialAddition(n int) int {
	for i := 0; i < n; i++ {
		return factorialAddition(n-1)
	}
	return n
}
