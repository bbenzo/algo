package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestSubArrayReplaceOnesLength1(t *testing.T) {
	arr := []int{1, 1, 0, 0, 1, 0, 1, 1, 0, 1}
	k := 2

	result := LongestSubArrayReplaceOnesLength(arr, k)

	assert.Equal(t, 6, result)
}

func TestLongestSubArrayReplaceOnesLength2(t *testing.T) {
	arr := []int{1, 1, 0, 0, 1, 0, 1, 1, 0, 1}
	k := 4

	result := LongestSubArrayReplaceOnesLength(arr, k)

	assert.Equal(t, len(arr), result)
}

func TestLongestSubArrayReplaceOnesLengthOnlyZeros(t *testing.T) {
	arr := []int{0, 0, 0, 0}
	k := 3

	result := LongestSubArrayReplaceOnesLength(arr, k)

	assert.Equal(t, 3, result)
}

func TestLongestSubArrayReplaceOnesLengthOnlyOnes(t *testing.T) {
	arr := []int{1, 1, 1, 1}
	k := 3

	result := LongestSubArrayReplaceOnesLength(arr, k)

	assert.Equal(t, 4, result)
}
