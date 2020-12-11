package two_pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDutchNationalFlag(t *testing.T) {
	input := []int{1, 0, 2, 1, 0}

	expected := []int{0, 0, 1, 1, 2}

	assert.Equal(t, expected, DutchNationalFlag(input))
}

func TestDutchNationalFlag2(t *testing.T) {
	input := []int{2, 2, 0, 1, 2, 0}

	expected := []int{0, 0, 1, 2, 2, 2}

	assert.Equal(t, expected, DutchNationalFlag(input))
}