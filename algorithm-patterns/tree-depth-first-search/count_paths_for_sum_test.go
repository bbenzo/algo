package tree_depth_first_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountAllPathsForSum(t *testing.T) {
	root := &node{
		val:   1,
		left:  &node{
			val:   7,
			left:  &node{
				val:   6,
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
				val:   3,
			},
		},
	}

	sum := 12

	expected := 3

 	assert.Equal(t, expected, CountAllPathsForSum(root, sum))
}
