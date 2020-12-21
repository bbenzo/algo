package merge_intervals

import "math"

func InsertInterval(intervals Intervals, interval []int) Intervals {
	result := Intervals{}

	// find place to insert in sorted list
	start, end := -1, -1
	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] < interval[0] || intervals[i][0] > interval[1] {
			result = append(result, intervals[i])
			continue
		}

		start = int(math.Min(float64(interval[0]), float64(intervals[i][0])))
		end = int(math.Max(float64(interval[1]), float64(intervals[i][1])))
	}

	result = append(result, []int{start, end})
	return result
}
