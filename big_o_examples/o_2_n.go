package big_o_examples

func recursiveFibonacci(num int) int {
	if num <= 1 {
		return num
	}
	return recursiveFibonacci(num - 2) + recursiveFibonacci(num - 1)
}
