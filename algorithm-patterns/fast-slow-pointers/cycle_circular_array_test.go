package fast_slow_pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCycleCircularArray(t *testing.T) {
	arr := []int{1, 2, -1, 2, 2}

	assert.True(t, CycleCircularArray(arr))
}

func TestCycleCircularArray2(t *testing.T) {
	arr := []int{2, 2, -1, 2}

	assert.True(t, CycleCircularArray(arr))
}

func TestCycleCircularArrayFalse(t *testing.T) {
	arr := []int{2, 1, -1, -2}

	assert.False(t, CycleCircularArray(arr))
}

