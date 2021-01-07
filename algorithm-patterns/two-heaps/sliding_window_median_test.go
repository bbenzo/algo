package two_heaps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlidingWindowMedian(t *testing.T) {
	nums := []int{1, 2, -1, 3, 5}
	k := 2

	expected := []float64{1.5, 0.5, 1.0, 4.0}

	assert.Equal(t, expected, SlidingWindowMedian(nums, k))
}
