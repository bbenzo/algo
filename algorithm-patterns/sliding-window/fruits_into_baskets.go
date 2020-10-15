package sliding_window

func FruitsIntoBaskets(fruits []string, baskets int) []string {
	start := 0
	used := make(map[string]int)
	var window []string
	var result []string

	for end := 0; end < len(fruits); end++ {
		window = append(window, fruits[end])
		used[fruits[end]] = end

		for len(used) > baskets {
			if used[fruits[start]] == start {
				delete(used, fruits[start])
				window = fruits[start+1:end+1]
			}

			start++
		}

		if len(window) > len(result) {
			result = window
		}
	}

	return result
}
