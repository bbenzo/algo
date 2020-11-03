package sliding_window

func ContainsPermutation(str, pattern string) bool {
	charMap := make(map[byte]bool, len(pattern))
	for i := range pattern {
		charMap[pattern[i]] = true
	}

	i, count := 0, 0
	for i < len(str) {
		if _, exists := charMap[str[i]]; exists {
			count++
		} else {
			count = 0
		}

		if count == len(pattern) {
			return true
		}

		i++
	}

	return false
}
