package modified_binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumDifference(t *testing.T) {
	input := []int{4, 6, 10}
	key := 7

	expected := 6

	assert.Equal(t, expected, MinimumDifference(input, key))
}

func TestMinimumDifference2(t *testing.T) {
	input := []int{4, 6, 10}
	key := 4

	expected := 4

	assert.Equal(t, expected, MinimumDifference(input, key))
}

func TestMinimumDifference3(t *testing.T) {
	input := []int{4, 6, 10}
	key := 17

	expected := 10

	assert.Equal(t, expected, MinimumDifference(input, key))
}

func TestMinimumDifference4(t *testing.T) {
	input := []int{1, 3, 8, 10, 15}
	key := 12

	expected := 10

	assert.Equal(t, expected, MinimumDifference(input, key))
}
