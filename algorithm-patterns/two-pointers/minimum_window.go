package two_pointers

import "sort"

func MinimumWindow(arr []int) int {
	cpy := make([]int, len(arr))
	copy(cpy, arr)

	sort.Ints(cpy)

	windowStart, windowEnd := -1, -1
	start, end := 0, len(arr)-1
	for start < end {
		if arr[start] != cpy[start] && windowStart == -1 {
			windowStart = start
		}
		start++

		if arr[end] != cpy[end] && windowEnd == -1 {
			windowEnd = end
		}
		end--
	}

	if windowEnd == -1 || windowStart == -1 {
		return 0
	}

	return len(arr[windowStart: windowEnd+1])
}
