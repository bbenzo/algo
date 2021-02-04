package modified_binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderAgnosticBinarySearch(t *testing.T) {
	input := []int{4, 6, 10}
	key := 10

	result := OrderAgnosticBinarySearch(input, key)

	assert.Equal(t, 2, result)
}

func TestOrderAgnosticBinarySearch2(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	key := 5

	result := OrderAgnosticBinarySearch(input, key)

	assert.Equal(t, 4, result)
}

func TestOrderAgnosticBinarySearch3(t *testing.T) {
	input := []int{10, 6, 4}
	key := 10

	result := OrderAgnosticBinarySearch(input, key)

	assert.Equal(t, 0, result)
}

func TestOrderAgnosticBinarySearch4(t *testing.T) {
	input := []int{4, 6, 10}
	key := 4

	result := OrderAgnosticBinarySearch(input, key)

	assert.Equal(t, 0, result)
}
