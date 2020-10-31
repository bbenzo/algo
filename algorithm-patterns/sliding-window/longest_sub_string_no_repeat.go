package sliding_window

func LongestSubStringNoRepeat(input string) string {
	end := 1
	window, result := []byte(input[:end]), []byte(input[:end])
	for end < len(input) {
		if input[end] != input[end - 1] {
			window = append(window, input[end])

			if len(window) > len(result) {
				result = window
			}
		} else {
			window = []byte{input[end]}
		}

		end++
	}

	return string(result)
}
