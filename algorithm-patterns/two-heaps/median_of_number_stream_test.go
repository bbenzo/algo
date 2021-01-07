package two_heaps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMedianNumberOfStream(t *testing.T) {
	tt := []struct {
		vals     []int
		expected float64
	}{
		{[]int{3, 1, 5}, 3},
		{[]int{1, 2, 3, 5, 4}, 3},
		{[]int{3, 5, 1, 9, 7}, 5},
	}

	for _, tc := range tt {
		s := newStream()

		for _, val := range tc.vals {
			s.insert(val)

		}
		assert.Equal(t, tc.expected, s.median())
	}
}
