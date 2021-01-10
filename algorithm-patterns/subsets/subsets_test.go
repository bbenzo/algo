package subsets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsets(t *testing.T) {
	input := []int{1, 3}
	expected := [][]int{{}, {1}, {3}, {1, 3}}

	assert.Equal(t, expected, Subsets(input))
}

func TestSubsets2(t *testing.T) {
	input := []int{1, 5, 3}
	expected := [][]int{{}, {1}, {5}, {1, 5}, {3}, {1, 3}, {5, 3}, {1, 5, 3}}

	assert.Equal(t, expected, Subsets(input))
}
