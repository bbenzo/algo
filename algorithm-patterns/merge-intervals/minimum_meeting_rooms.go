package merge_intervals

import "sort"

func MinimumMeetingRooms(intervals Intervals) int {
	sort.Sort(&intervals)

	count := 1
	run(intervals, &count)

	return count
}

func run(intervals Intervals, count *int) {
	var intersections Intervals

	for i := 1; i < len(intervals); i++ {
		if intervals[i-1][1] > intervals[i][0] {
			intersections = append(intersections, []int{intervals[i][0], intervals[i-1][1]})
		}
	}

	if len(intersections) > 0 {
		*count++
	}

	if len(intersections) > 1 {
		run(intersections, count)
	}
}
