package modified_binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotationCount(t *testing.T) {
	input := []int{10, 15, 1, 3, 8}

	expected := 2

	assert.Equal(t, expected, RotationCount(input))
}

func TestRotationCount2(t *testing.T) {
	input := []int{4, 5, 7, 9, 10, -1, 2}

	expected := 5

	assert.Equal(t, expected, RotationCount(input))
}

func TestRotationCount3(t *testing.T) {
	input := []int{1, 3, 8, 10}

	expected := 0

	assert.Equal(t, expected, RotationCount(input))
}
