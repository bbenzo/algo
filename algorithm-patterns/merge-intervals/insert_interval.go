package merge_intervals

import (
	"sort"
)

func InsertInterval(intervals Intervals, interval []int) Intervals {
	sort.Sort(&intervals)

	result := Intervals{}

	mergedWithInterval := []int{interval[0], interval[1]}
	for i := 0; i < len(intervals); i++ {
		current := intervals[i]

		if (current[0] >= interval[0] && current[0] <= interval[1]) ||
			(current[1] >= interval[0] && current[0] <= interval[1]) {
			mergedWithInterval = append(mergedWithInterval, current...)
		} else {
			result = append(result, current)
		}
	}

	sort.Ints(mergedWithInterval)

	result = append(result, []int{mergedWithInterval[0], mergedWithInterval[len(mergedWithInterval) - 1]})

	return result
}
