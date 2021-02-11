package bitwise_xor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlipInvertImage(t *testing.T) {
	input := [][]int{
		{1, 0, 1},
		{1, 1, 1},
		{0, 1, 1},
	}

	output := [][]int{
		{0, 1, 0},
		{0, 0, 0},
		{0, 0, 1},
	}

	assert.Equal(t, output, FlipInvertImage(input))
}
