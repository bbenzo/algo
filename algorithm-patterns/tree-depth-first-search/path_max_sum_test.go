package tree_depth_first_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPathMaxSum(t *testing.T) {
	root := &node{
		val:   1,
		left:  &node{
			val:   2,
			left: &node{
				val:   4,
			},
		},
		right: &node{
			val:   3,
			left:  &node{
				val:   5,
			},
			right: &node{
				val:   6,
			},
		},
	}

	expected := []int{4, 2, 1, 3, 6}

	path, sum := PathMaxSum(root)

 	assert.Equal(t, expected, path)
 	assert.Equal(t, 16, sum)
}
