package merge_intervals

import "sort"

func ConflictingAppointments(intervals Intervals) bool {
	sort.Sort(&intervals)

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}
