package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	given := []int{4, 1, 23, 34, 20, 11, 40}
	expected := []int{1, 4, 11, 20, 23, 34, 40}

	result := MergeSort(given)
	assert.Equal(t, expected, result)
}

func TestMergeSortOneItem(t *testing.T) {
	given := []int{23}
	expected := []int{23}

	result := MergeSort(given)
	assert.Equal(t, expected, result)
}

func TestMergeSortNoItem(t *testing.T) {
	given := []int{}
	expected := []int{}

	result := MergeSort(given)
	assert.Equal(t, expected, result)
}
