package subsets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsetsWithDuplicates(t *testing.T) {
	input := []int{1, 3, 3}
	expected := [][]int{{}, {1}, {3}, {1, 3}, {3, 3}, {1, 3, 3}}

	assert.Equal(t, expected, SubsetsWithDuplicates(input))
}

func TestSubsetsWithDuplicates2(t *testing.T) {
	input := []int{1, 5, 3, 3}
	expected := [][]int{{}, {1}, {3}, {1, 3}, {3, 3}, {1, 3, 3}, {5}, {1, 5}, {3, 5}, {1, 3, 5}, {3, 3, 5}, {1, 3, 3, 5}}

	assert.Equal(t, expected, SubsetsWithDuplicates(input))
}
