package tree_depth_first_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllPathsForSum(t *testing.T) {
	root := &node{
		val:   1,
		left:  &node{
			val:   7,
			left:  &node{
				val:   4,
			},
			right: &node{
				val:   5,
			},
		},
		right: &node{
			val:   9,
			left:  &node{
				val:   2,
			},
			right: &node{
				val:   7,
			},
		},
	}

	sum := 12

	expected := [][]int{{1, 7, 4}, {1, 9, 2}}

 	assert.Equal(t, expected, AllPathsForSum(root, sum))
}
