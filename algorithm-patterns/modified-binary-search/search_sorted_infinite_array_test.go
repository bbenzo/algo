package modified_binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchSortedInfiniteArray(t *testing.T) {
	input := []int{4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30}
	key := 16

	expected := 6

	assert.Equal(t, expected, SearchSortedInfiniteArray(input, key))
}
