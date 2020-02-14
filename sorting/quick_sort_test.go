package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSortRandom(t *testing.T) {
	given := []int{4, 1, 23, 34, 20, 11, 40}
	expected := []int{1, 4, 11, 20, 23, 34, 40}

	result := QuickSort(given)
	assert.Equal(t, expected, result)
}
