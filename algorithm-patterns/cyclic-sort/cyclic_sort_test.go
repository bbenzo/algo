package cyclic_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCyclicSort1(t *testing.T) {
	arr := []int{3, 1, 5, 4, 2}
	expected := []int{1, 2, 3, 4, 5}

	CyclicSort(arr)

	assert.Equal(t, expected, arr)
}
