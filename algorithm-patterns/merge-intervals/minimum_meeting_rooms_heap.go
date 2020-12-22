package merge_intervals

import "sort"

type Meetings []Meeting

type Meeting struct {
	start int
	end int
}

func (m *Meetings) Len() int {
	return len(*m)
}

func (m *Meetings) Less(i, j int) bool {
	return (*m)[i].end < (*m)[j].end
}

func (m *Meetings) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

type MinHeap []Meeting

func (m *MinHeap) Len() int {
	return len(*m)
}

func (m *MinHeap) Less(i, j int) bool {
	return (*m)[i].end < (*m)[j].end
}

func (m *MinHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *MinHeap) Push(x Meeting) {
	*m = append(*m, x)
}

func (m *MinHeap) Pop() Meeting {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

func (m *MinHeap) Peek() Meeting {
	return (*m)[0]
}

func MinimumMeetingRoomsHeap(meetings Meetings) int {
	sort.Sort(&meetings)

	minRooms := 0
	minHeap := MinHeap{}
	for _, meeting := range meetings {
		for minHeap.Len() > 0 && meeting.start >= minHeap.Peek().end {
			minHeap.Pop()
		}

		minHeap.Push(meeting)

		minRooms = max(minRooms, minHeap.Len())
	}

	return minRooms
}
