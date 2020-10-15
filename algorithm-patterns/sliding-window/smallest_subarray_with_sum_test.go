package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfSmallestContiguousSubArrayWithSum(t *testing.T) {
	arr := []int{9, 1, 7}
	result := LengthOfSmallestContiguousSubArrayWithSum(arr, 10)

	assert.Equal(t, 2, result)
}

func TestLengthOfSmallestContiguousSubArrayWithSum2(t *testing.T) {
	arr := []int{7, 1, 1, 4, 5, 1}
	result := LengthOfSmallestContiguousSubArrayWithSum(arr, 10)

	assert.Equal(t, 3, result)
}

func TestLengthOfSmallestContiguousSubArrayWithSum3(t *testing.T) {
	arr := []int{4, 5, 1, 7, 1, 1}
	result := LengthOfSmallestContiguousSubArrayWithSum(arr, 30)

	assert.Equal(t, 0, result)
}

func TestLengthOfSmallestContiguousSubArrayWithSum4(t *testing.T) {
	arr := []int{4, 5, 1, 10, 1}
	result := LengthOfSmallestContiguousSubArrayWithSum(arr, 10)

	assert.Equal(t, 1, result)
}
