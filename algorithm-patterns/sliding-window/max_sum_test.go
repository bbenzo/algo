package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSumOfSubArray(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 1}

	result := MaxSumOfSubArray(arr, 3)
	assert.Equal(t, 21, result)
}