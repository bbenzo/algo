package sliding_window

func StringAnagramIndexes(str string, pattern string) []int {
	start, end := 0, 0
	charMap := make(map[byte]int)
	for i := range pattern {
		charMap[pattern[i]]++
	}

	var result []int

	windowCharMap := make(map[byte]int)
	correctMap := make(map[byte]bool)
	for end < len(str) {
		windowCharMap[str[end]]++

		if end >= len(pattern) {
			windowCharMap[str[start]]--
			if windowCharMap[str[start]] == charMap[str[start]] {
				correctMap[str[start]] = true
			} else {
				delete(correctMap, str[start])
			}

			start++
		}

		if windowCharMap[str[end]] == charMap[str[end]] {
			correctMap[str[end]] = true
		} else {
			delete(correctMap, str[end])
		}

		if len(correctMap) == len(charMap) {
			result = append(result, start)
		}

		end++
	}

	return result
}
