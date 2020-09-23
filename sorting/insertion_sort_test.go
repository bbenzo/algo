package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	given := []int{1, 4, 11, 20, 23, 34, 40}
	expected := []int{1, 4, 5, 11, 20, 23, 34, 40}

	result := Insert(given, 5)
	assert.Equal(t, expected, result)
}

func TestInsertEmptyArr(t *testing.T) {
	given := []int{}
	expected := []int{5}

	result := Insert(given, 5)
	assert.Equal(t, expected, result)
}

func TestInsertionSort(t *testing.T) {
	given := []int{4, 1, 23, 34, 20, 11, 40}
	expected := []int{1, 4, 11, 20, 23, 34, 40}

	result := InsertionSort(given)
	assert.Equal(t, expected, result)
}

func TestInsertionSortOneItem(t *testing.T) {
	given := []int{4}
	expected := []int{4}

	result := InsertionSort(given)
	assert.Equal(t, expected, result)
}

func TestInsertionSortNoItem(t *testing.T) {
	given := []int{}
	expected := []int{}

	result := InsertionSort(given)
	assert.Equal(t, expected, result)
}
