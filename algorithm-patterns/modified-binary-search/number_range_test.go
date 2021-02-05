package modified_binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberRange(t *testing.T) {
	input := []int{4, 6, 6, 6, 9}
	key := 6

	expected := []int{1, 3}

	assert.Equal(t, expected, NumberRange(input, key))
}

func TestNumberRange2(t *testing.T) {
	input := []int{1, 3, 8, 10, 15}
	key := 10

	expected := []int{3, 3}

	assert.Equal(t, expected, NumberRange(input, key))
}

func TestNumberRange3(t *testing.T) {
	input := []int{1, 3, 8, 10, 15}
	key := 12

	expected := []int{-1, -1}

	assert.Equal(t, expected, NumberRange(input, key))
}
