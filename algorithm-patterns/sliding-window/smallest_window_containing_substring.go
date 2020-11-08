package sliding_window

func SmallestStringContainingPattern(str, pattern string) string {
	start, end := 0, 0
	var result []byte
	var window []byte
	charCounts := make(map[byte]int)
	charExists := make(map[byte]bool, 3)
	for i := range pattern {
		charExists[pattern[i]] = true
	}

	for end < len(str) {
		if charExists[str[end]] {
			charCounts[str[end]]++
		}

		window = append(window, str[end])

		for len(charCounts) >= len(pattern) {
			charCounts[str[start]]--

			if charCounts[str[start]] < 1 {
				window = []byte(str[start:end+1])
				if len(window) < len(result) || len(result) == 0 {
					result = window
				}

				delete(charCounts, str[start])
			}

			start++
		}

		end++
	}

	return string(result)
}
