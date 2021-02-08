package modified_binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitonicArrayMax(t *testing.T) {
	input := []int{1, 3, 8, 12, 4, 2}

	expected := 12

	assert.Equal(t, expected, BitonicArrayMax(input))
}

func TestBitonicArrayMax2(t *testing.T) {
	input := []int{3, 8, 3, 1}

	expected := 8

	assert.Equal(t, expected, BitonicArrayMax(input))
}

func TestBitonicArrayMax3(t *testing.T) {
	input := []int{1, 3, 8, 12}

	expected := 12

	assert.Equal(t, expected, BitonicArrayMax(input))
}

func TestBitonicArrayMax4(t *testing.T) {
	input := []int{10, 9, 8}

	expected := 10

	assert.Equal(t, expected, BitonicArrayMax(input))
}
