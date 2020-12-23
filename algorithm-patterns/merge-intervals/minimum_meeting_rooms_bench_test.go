package merge_intervals

import (
	"testing"
)

func BenchmarkRecursive(b *testing.B) {
	meetings := Intervals{{4, 5}, {2, 3}, {2, 4}, {3, 5}, {3,5}}

	for i := 0; i < b.N; i++ {
		MinimumMeetingRooms(meetings)
	}
}

func BenchmarkHeap(b *testing.B) {
	meetings := Meetings{{4, 5}, {2, 3}, {2, 4}, {3, 5}, {3,5}}

	for i := 0; i < b.N; i++ {
		MinimumMeetingRoomsHeap(meetings)
	}
}


//goos: darwin
//goarch: amd64
//pkg: github/bbenzo/algo/algorithm-patterns/merge-intervals
//BenchmarkRecursive      18076503               653 ns/op             384 B/op         12 allocs/op
//BenchmarkHeap           58458181               202 ns/op             144 B/op          4 allocs/op