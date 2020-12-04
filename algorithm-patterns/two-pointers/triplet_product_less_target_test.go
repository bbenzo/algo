package two_pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTripletsWithProductLessTarget(t *testing.T) {
	input := []int{2, 5, 3, 10}
	target := 30

	expected := [][]int{{2}, {2, 5}, {5}, {5, 3}, {3}, {10}}

	assert.Equal(t, expected, TriplestWithProductLessTarget(input, target))
}

func TestTripletsWithProductLessTarget2(t *testing.T) {
	input := []int{8, 2, 6, 5}
	target := 50

	expected := [][]int{{8}, {8, 2}, {2}, {2, 6}, {6}, {6, 5}, {5}}

	assert.Equal(t, expected, TriplestWithProductLessTarget(input, target))
}
