package bitwise_xor

func MissingNumber(arr []int) int {
	n := len(arr) + 1

	x1 := 1
	for i := 2; i <= n; i++ {
		x1 = x1 ^ i
	}

	x2 := arr[0]
	for i := 1; i < n-1; i++ {
		x2 = x2 ^ arr[i]
	}

	return x1 ^ x2
}
