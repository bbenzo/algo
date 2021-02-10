package bitwise_xor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	arr := []int{1, 4, 2, 1, 3, 2, 3}

	assert.Equal(t, 4, SingleNumber(arr))
}

func TestSingleNumber2(t *testing.T) {
	arr := []int{7, 9, 7}

	assert.Equal(t, 9, SingleNumber(arr))
}
