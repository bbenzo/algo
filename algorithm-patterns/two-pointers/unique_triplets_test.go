package two_pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniqueTriplets(t *testing.T) {
	input := []int{-3, 0, 1, 2, -1, 1, -2} // sorted = -3, -2, -1, 0, 1, 1, 2
	expected := [][]int{{-3, 2, 1}, {-2, 2, 0}, {-2, 1, 1}, {-1, 1, 0}}

	result := UniqueTripletsThatAddUpToZero(input)

	assert.Equal(t, expected, result)
}
