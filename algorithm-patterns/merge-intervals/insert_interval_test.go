package merge_intervals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertInterval(t *testing.T) {
	intervals := Intervals{{1, 3}, {5, 7}, {8, 12}}
	interval := []int{4, 6}

	expected := Intervals{{1, 3}, {8, 12}, {4, 7}}

	assert.Equal(t, expected, InsertInterval(intervals, interval))
}

func TestInsertInterval2(t *testing.T) {
	intervals := Intervals{{1, 3}, {5, 7}, {8, 12}}
	interval := []int{4, 10}

	expected := Intervals{{1, 3}, {4, 12}}

	assert.Equal(t, expected, InsertInterval(intervals, interval))
}

func TestInsertInterval3(t *testing.T) {
	intervals := Intervals{{2, 3}, {5, 7}}
	interval := []int{1, 4}

	expected := Intervals{{5, 7}, {1, 4}}

	assert.Equal(t, expected, InsertInterval(intervals, interval))
}
