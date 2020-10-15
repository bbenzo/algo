package sliding_window

// Given a string, find the length of the longest substring in it with no more than K distinct characters.
func LongestSubStringWithKDistinctCharacters(word string, k int) string {
	start := 0                 // first element of the observed window
	used := make(map[byte]int) // map with all bytes and their last appearance in the word
	var window []byte          // the observed window
	var result []byte          // the longest substring

	for end := 0; end < len(word); end++ {
		window = append(window, word[end])
		used[word[end]] = end

		for len(used) > k {
			if used[word[start]] == start {
				delete(used, word[start])
				window = []byte(word[start+1 : end+1])
			}

			start++
		}

		if len(window) > len(result) {
			result = window
		}
	}

	if len(result) < k {
		return ""
	}

	return string(result)
}
