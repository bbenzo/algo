package two_pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumOfPairEqualSum(t *testing.T) {
	arr := []int{1, 4, 6, 8, 10, 12}
	targetSum := 10

	index1, index2 := SumOfPairEqualSum(arr, targetSum)

	assert.Equal(t, 1, index1)
	assert.Equal(t, 2, index2)
}

func TestSumOfPairEqualSum2(t *testing.T) {
	arr := []int{3, 4, 5, 11, 15}
	targetSum := 16

	index1, index2 := SumOfPairEqualSum(arr, targetSum)

	assert.Equal(t, 2, index1)
	assert.Equal(t, 3, index2)
}

func TestSumOfPairEqualSum3(t *testing.T) {
	arr := []int{6, 10, 13, 24, 56}
	targetSum := 80

	index1, index2 := SumOfPairEqualSum(arr, targetSum)

	assert.Equal(t, 3, index1)
	assert.Equal(t, 4, index2)
}
