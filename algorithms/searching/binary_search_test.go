package searching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	list := []int{2, 4, 6, 8, 10, 12, 14, 16}
	search := 4
	result := BinarySearch(search, list)

	assert.Equal(t, 1, result)
}

func TestBinarySearchNotPresent(t *testing.T) {
	list := []int{2, 4, 6, 8, 10, 12, 14, 16}
	search := 3
	result := BinarySearch(search, list)

	assert.Equal(t, -1, result)
}
