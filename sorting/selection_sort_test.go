package sorting

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestSelectionSortWorstCase(t *testing.T) {
	sorted := []int32{1, 2, 3, 4, 5}

	list := []int32{5, 4, 3, 2, 1}

	result:= SelectionSort(list)

	assert.Equal(t, result, sorted)
}
