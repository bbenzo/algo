package cyclic_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMissingNums1(t *testing.T) {
	arr := []int{2, 3, 1, 8, 2, 3, 5, 1}
	expected := []int{4, 6, 7}

	assert.Equal(t, expected, FindMissingNums(arr))
}

func TestFindMissingNums2(t *testing.T) {
	arr := []int{2, 4, 1, 2}
	expected := []int{3}

	assert.Equal(t, expected, FindMissingNums(arr))
}

func TestFindMissingNums3(t *testing.T) {
	arr := []int{2, 3, 2, 1}
	expected := []int{4}

	assert.Equal(t, expected, FindMissingNums(arr))
}
