package merge_intervals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumMeetingRoomsHeap(t *testing.T) {
	meetings := Meetings{{1, 4}, {2, 5}, {7, 9}}

	assert.Equal(t, 2, MinimumMeetingRoomsHeap(meetings))
}

func TestMinimumMeetingRoomsHeap2(t *testing.T) {
	meetings := Meetings{{6, 7}, {2, 4}, {8, 12}}

	assert.Equal(t, 1, MinimumMeetingRoomsHeap(meetings))
}

func TestMinimumMeetingRoomsHeap3(t *testing.T) {
	meetings := Meetings{{1, 4}, {2, 3}, {3, 6}}

	assert.Equal(t, 2, MinimumMeetingRoomsHeap(meetings))
}

func TestMinimumMeetingRoomsHeap4(t *testing.T) {
	meetings := Meetings{{4, 5}, {2, 3}, {2, 4}, {3, 5}}

	assert.Equal(t, 2, MinimumMeetingRoomsHeap(meetings))
}

func TestMinimumMeetingRoomsHeap5(t *testing.T) {
	meetings := Meetings{{4, 5}, {2, 3}, {2, 4}, {3, 5}, {3,5}}

	assert.Equal(t, 3, MinimumMeetingRoomsHeap(meetings))
}
