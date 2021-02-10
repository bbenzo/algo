package bitwise_xor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSingleNumbers(t *testing.T) {
	arr := []int{1, 4, 2, 1, 3, 5, 6, 2, 3, 5}

	assert.Equal(t, []int{6, 4}, TwoSingleNumbers(arr))
}

func TestTwoSingleNumbers2(t *testing.T) {
	arr := []int{2, 1, 3, 2}

	assert.Equal(t, []int{3, 1}, TwoSingleNumbers(arr))
}

func TestTwoSingleNumbers3(t *testing.T) {
	arr := []int{1, 1, 3, 4, 2, 2}

	assert.Equal(t, []int{3, 4}, TwoSingleNumbers(arr))
}
