package merge_intervals

import "sort"

func EmployeeFreeTime(employeeMeetings []Meetings) Meetings {
	allMeetings := Meetings{}
	for i := range employeeMeetings {
		allMeetings = append(allMeetings, employeeMeetings[i]...)
	}

	sort.Sort(&allMeetings)

	freeMeetings := Meetings{}
	for i := 1; i < len(allMeetings); i++ {
		if allMeetings[i].start - allMeetings[i-1].end > 0 {
			freeMeetings = append(freeMeetings, Meeting{allMeetings[i-1].end, allMeetings[i].start})
		}
	}

	return freeMeetings
}
