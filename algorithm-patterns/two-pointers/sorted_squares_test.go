package two_pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedSquares1(t *testing.T) {
	input := []int{-2, -1, 0, 2, 3}
	expected := []int{0, 1, 4, 4, 9}

	result := SortedSquares(input)

	assert.Equal(t, expected, result)
}

func TestSortedSquares2(t *testing.T) {
	input := []int{-2, -1, 0, 2, 3}
	expected := []int{0, 1, 4, 4, 9}

	result := SortedSquares(input)

	assert.Equal(t, expected, result)
}