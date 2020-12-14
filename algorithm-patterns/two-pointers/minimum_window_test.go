package two_pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumWindow1(t *testing.T) {
	input := []int{1, 2, 5, 3, 7, 10, 9, 12}

	assert.Equal(t, 5, MinimumWindow(input))
}

func TestMinimumWindow2(t *testing.T) {
	input := []int{1, 3, 2, 0, -1, 7, 10}

	assert.Equal(t, 5, MinimumWindow(input))
}

func TestMinimumWindow3(t *testing.T) {
	input := []int{1, 2, 3}

	assert.Equal(t, 0, MinimumWindow(input))
}

func TestMinimumWindow4(t *testing.T) {
	input := []int{3, 2, 1}

	assert.Equal(t, 3, MinimumWindow(input))
}
