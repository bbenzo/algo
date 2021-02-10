package bitwise_xor

func SingleNumber(arr []int) int {
	x := arr[0]
	for i := 1; i < len(arr); i++ {
		x = x ^ arr[i]
	}

	return x
}
