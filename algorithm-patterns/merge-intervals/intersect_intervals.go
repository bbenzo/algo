package merge_intervals

func IntersectIntervals(arr1, arr2 Intervals) Intervals {
	var result Intervals

	idx1, idx2 := 0, 0
	for idx1 < len(arr1) && idx2 < len(arr2) {
		interval1 := arr1[idx1]
		interval2 := arr2[idx2]

		if interval2[0] > interval1[1] {
			idx1++
			continue
		}

		intersectStart := max(interval1[0], interval2[0])
		intersectEnd := min(interval1[1], interval2[1])
		result = append(result, []int{intersectStart, intersectEnd})

		if idx2+1 < len(arr2) {
			idx2++
		} else {
			idx1++
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
