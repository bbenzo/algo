package sliding_window

func LongestSubStringReplace(input string, k int) string {
	start, end := 0, 0
	replaced := 0
	letterCount := make(map[byte]int)
	var window []byte
	var result []byte

	for end < len(input) {
		// count each letter
		letterCount[input[end]]++

		// append to window
		window = append(window, input[end])

		// check if this a letter which needs to be replaced
		if input[end] != input[start] {
			replaced++
		}

		// if we replaced more than k
		for replaced > k {
			// adapt letter count
			letterCount[input[start]]--

			// shift start
			start++

			// if we arrive at next letter which is not our
			// start letter, we set this as new start
			if input[start] != window[0] {
				// subtract all counts of the new letter from replaced count
				// and add remaining count of former start to replaced count
				replaced -= letterCount[input[start]]
				replaced += letterCount[window[0]]

				window = []byte(input[start:end+1])
			}
		}

		if len(window) > len(result) {
			result = window
		}

		end++
	}

	return string(result)
}
