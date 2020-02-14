package sorting

func merge(left []int, right []int) []int {
	merged := make([]int, len(left) + len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			merged[i] = left[0]
			left = left[1:]
		} else {
			merged[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
        merged[i] = left[j]
        i++
    }
    for j := 0; j < len(right); j++ {
        merged[i] = right[j]
        i++
    }

	return merged
}
