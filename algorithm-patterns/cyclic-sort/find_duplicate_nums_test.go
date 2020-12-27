package cyclic_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDuplicateNums1(t *testing.T) {
	arr := []int{3, 4, 4, 5, 5}

	expected := []int{5, 4}

	assert.Equal(t, expected, FindDuplicateNums(arr))
}

func TestFindDuplicateNums2(t *testing.T) {
	arr := []int{5, 4, 7, 2, 3, 5, 3}

	expected := []int{3, 5}

	assert.Equal(t, expected, FindDuplicateNums(arr))
}
