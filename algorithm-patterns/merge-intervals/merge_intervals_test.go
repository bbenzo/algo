package merge_intervals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	input := Intervals{{1, 4}, {2, 5}, {7, 9}}
	output := Intervals{{1, 5}, {7, 9}}

	assert.Equal(t, output, MergeIntervals(input))
}

func TestMergeIntervals2(t *testing.T) {
	input := Intervals{{6,7}, {2,4}, {5,9}}
	output := Intervals{{2,4}, {5,9}}

	assert.Equal(t, output, MergeIntervals(input))
}

func TestMergeIntervals3(t *testing.T) {
	input := Intervals{{1, 4}, {2, 6}, {3,5}}
	output := Intervals{{1, 6}}

	assert.Equal(t, output, MergeIntervals(input))
}
