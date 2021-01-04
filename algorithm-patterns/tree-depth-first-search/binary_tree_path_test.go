package tree_depth_first_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreePath(t *testing.T) {
	root := &node{
		val:   1,
		left:  &node{
			val:   2,
			left:  &node{
				val:   4,
			},
			right: &node{
				val:   5,
			},
		},
		right: &node{
			val:   3,
			left:  &node{
				val:   6,
			},
			right: &node{
				val:   7,
			},
		},
	}

	assert.True(t, BinaryTreePath(root, 11))
}

func TestBinaryTreePathFalse(t *testing.T) {
	root := &node{
		val:   1,
		left:  &node{
			val:   2,
			left:  &node{
				val:   4,
			},
			right: &node{
				val:   5,
			},
		},
		right: &node{
			val:   3,
			left:  &node{
				val:   6,
			},
			right: &node{
				val:   7,
			},
		},
	}

	assert.False(t, BinaryTreePath(root, 30))
}
