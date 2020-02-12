package big_o_examples

func factorialAddition(n int) int {
	for i := 0; i < n; i++ {
		return factorialAddition(n-1)
	}
	return n
}
