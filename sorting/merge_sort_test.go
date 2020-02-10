package sorting

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	given := []int{4, 1, 23, 34, 20, 11, 40}
	expected := []int{1, 4, 11, 20, 23, 34, 40}

	result := MergeSort(given)
	assert.Equal(t, expected, result)
}
