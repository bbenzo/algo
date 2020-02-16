package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectionSortWorstCase(t *testing.T) {
	given := []int32{5, 4, 3, 2, 1}
	expected := []int32{1, 2, 3, 4, 5}

	result := SelectionSort(given)

	assert.Equal(t, result, expected)
}

func TestSelectionSortWorstCaseOneItem(t *testing.T) {
	given := []int32{1}
	expected := []int32{1}

	result := SelectionSort(given)

	assert.Equal(t, result, expected)
}

func TestSelectionSortWorstCaseNoItem(t *testing.T) {
	given := []int32{}
	expected := []int32{}

	result := SelectionSort(given)

	assert.Equal(t, result, expected)
}
