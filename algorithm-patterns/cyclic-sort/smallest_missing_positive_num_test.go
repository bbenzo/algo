package cyclic_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSmallestMissingPositiveNum(t *testing.T) {
	arr := []int{-3, 1, 5, 4, 2}

	assert.Equal(t, 3, SmallestMissingPositiveNum(arr))
}

func TestSmallestMissingPositiveNum2(t *testing.T) {
	arr := []int{3, -2, 0, 1, 2}

	assert.Equal(t, 4, SmallestMissingPositiveNum(arr))
}

func TestSmallestMissingPositiveNum3(t *testing.T) {
	arr := []int{3, 2, 5, 1}

	assert.Equal(t, 4, SmallestMissingPositiveNum(arr))
}