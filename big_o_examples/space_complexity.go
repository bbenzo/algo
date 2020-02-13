package big_o_examples

// n is the length of the array
func sum(arr []int, n int) int {
	var result int // 4 bytes for result
	for i := 0; i < n; i++ { // 4 bytes for i
		result += arr[i]
	}
	return result
}
