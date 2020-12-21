package merge_intervals

import "sort"

type Intervals [][]int

func (in *Intervals) Len() int {
	return len(*in)
}

func (in *Intervals) Less(i, j int) bool {
	return (*in)[i][0] < (*in)[j][0]
}

func (in *Intervals) Swap(i, j int) {
	(*in)[i], (*in)[j] = (*in)[j], (*in)[i]
}

func MergeIntervals(intervals Intervals) Intervals {
	sort.Sort(&intervals)

	result := [][]int{intervals[0]}

	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]

		if interval[0] <= end {
			if interval[1] > end {
				end = interval[1]
				result[i-1][1] = end
			}
			continue
		} else {
			start, end = interval[0], interval[1]
			result = append(result, []int{start, end})
		}
	}

	return result
}
