package sliding_window

func FindAveragesOfSubArrays(arr []int, length int) []float64 {
	var result []float64

	if length > len(arr) {
		return result
	}

	windowStart, sum := 0, 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]

		if i >= length - 1 {
			result = append(result, float64(sum) / float64(length))
			sum -= arr[windowStart]
			windowStart++
		}
	}

	return result
}
