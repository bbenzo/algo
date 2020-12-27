package cyclic_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSmallestMissingKPositiveNums(t *testing.T) {
	arr := []int{3, -1, 4, 5, 5}
	k := 3

	expected := []int{1, 2, 6}

	assert.Equal(t, expected, SmallestMissingKPositiveNums(arr, k))
}

func TestSmallestMissingKPositiveNums2(t *testing.T) {
	arr := []int{2, 3, 4}
	k := 3

	expected := []int{1, 5, 6}

	assert.Equal(t, expected, SmallestMissingKPositiveNums(arr, k))
}

func TestSmallestMissingKPositiveNums3(t *testing.T) {
	arr := []int{-2, -3, 4}
	k := 2

	expected := []int{1, 2}

	assert.Equal(t, expected, SmallestMissingKPositiveNums(arr, k))
}
