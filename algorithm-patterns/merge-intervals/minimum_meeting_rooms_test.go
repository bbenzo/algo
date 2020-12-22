package merge_intervals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumMeetingRooms(t *testing.T) {
	meetings := Intervals{{1, 4}, {2, 5}, {7, 9}}

	assert.Equal(t, 2, MinimumMeetingRooms(meetings))
}

func TestMinimumMeetingRooms2(t *testing.T) {
	meetings := Intervals{{6, 7}, {2, 4}, {8, 12}}

	assert.Equal(t, 1, MinimumMeetingRooms(meetings))
}

func TestMinimumMeetingRooms3(t *testing.T) {
	meetings := Intervals{{1, 4}, {2, 3}, {3, 6}}

	assert.Equal(t, 2, MinimumMeetingRooms(meetings))
}

func TestMinimumMeetingRooms4(t *testing.T) {
	meetings := Intervals{{4, 5}, {2, 3}, {2, 4}, {3, 5}}

	assert.Equal(t, 2, MinimumMeetingRooms(meetings))
}

func TestMinimumMeetingRooms5(t *testing.T) {
	meetings := Intervals{{4, 5}, {2, 3}, {2, 4}, {3, 5}, {3,5}}

	assert.Equal(t, 3, MinimumMeetingRooms(meetings))
}
