package modified_binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchRotatedArray(t *testing.T) {
	input := []int{10, 15, 1, 3, 8}
	key := 15

	expected := 1

	assert.Equal(t, expected, SearchRotatedArray(input, key))
}

func TestSearchRotatedArray2(t *testing.T) {
	input := []int{4, 5, 7, 9, 10, -1, 2}
	key := 10

	expected := 4

	assert.Equal(t, expected, SearchRotatedArray(input, key))
}
