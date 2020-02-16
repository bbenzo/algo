package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	given := []int{4, 1, 23, 34, 50, 20, 11,}
	expected := []int{1, 4, 11, 20, 23, 34, 50}

	result := QuickSort(given)
	assert.Equal(t, expected, result)
}

func TestQuickSortArrayWithOneElement(t *testing.T) {
	given := []int{4}
	expected := []int{4}

	result := QuickSort(given)
	assert.Equal(t, expected, result)
}

func TestQuickSortArrayWithNoElement(t *testing.T) {
	given := []int{}
	expected := []int{}

	result := QuickSort(given)
	assert.Equal(t, expected, result)
}
