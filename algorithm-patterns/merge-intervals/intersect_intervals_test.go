package merge_intervals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersectIntervals(t *testing.T) {
	arr1 := Intervals{{1, 3}, {5, 6}, {7, 9}}
	arr2 := Intervals{{2, 3}, {5, 7}}

	expected := Intervals{{2, 3}, {5, 6}, {7, 7}}

	assert.Equal(t, expected, IntersectIntervals(arr1, arr2))
}

func TestIntersectIntervals2(t *testing.T) {
	arr1 := Intervals{{1, 3}, {5, 7}, {9, 12}}
	arr2 := Intervals{{5, 10}}

	expected := Intervals{{5, 7}, {9, 10}}

	assert.Equal(t, expected, IntersectIntervals(arr1, arr2))
}
