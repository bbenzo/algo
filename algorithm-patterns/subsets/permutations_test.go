package subsets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermutations(t *testing.T) {
	input := []int{1, 3, 5}
	expected := [][]int{[]int{1, 3, 5}, []int{1, 5, 3}, []int{5, 1, 3}, []int{3, 1, 5}, []int{3, 5, 1}, []int{5, 3, 1}}

	assert.Equal(t, expected, Permutations(input))
}
