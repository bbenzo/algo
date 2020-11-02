package sliding_window

func LongestSubArrayReplaceOnesLength(arr []int, k int) int {
	start, end := 0, 0
	replaced := 0
	var window []int
	var result []int

	for end < len(arr) {
		window = append(window, arr[end])

		if arr[end] != 1 {
			replaced++
		}

		for replaced > k {
			if arr[start] != 1 {
				replaced--
			}

			start++
			window = arr[start:end+1]
		}

		if len(window) > len(result) {
			result = window
		}

		end++
	}

	return len(result)
}
