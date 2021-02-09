package modified_binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchBitonicArray(t *testing.T) {
	input := []int{1, 3, 8, 4, 3}
	key := 4

	expected := 3

	assert.Equal(t, expected, SearchBitonicArray(input, key))
}

func TestSearchBitonicArray2(t *testing.T) {
	input := []int{3, 8, 3, 1}
	key := 8

	expected := 1

	assert.Equal(t, expected, SearchBitonicArray(input, key))
}

func TestSearchBitonicArray3(t *testing.T) {
	input := []int{1, 3, 8, 12}
	key := 12

	expected := 3

	assert.Equal(t, expected, SearchBitonicArray(input, key))
}

func TestSearchBitonicArray4(t *testing.T) {
	input := []int{10, 9, 8}
	key := 10

	expected := 0

	assert.Equal(t, expected, SearchBitonicArray(input, key))
}
