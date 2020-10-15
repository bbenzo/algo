package sliding_window

func MaxSumOfSubArray(arr []int, length int) int {
	if length > len(arr) {
		length = len(arr)
	}

	start, sum, max := 0, 0, 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if i >= length - 1 {
			if sum > max {
				max = sum
			}

			sum -= arr[start]
			start++
		}
	}

	return max
}
