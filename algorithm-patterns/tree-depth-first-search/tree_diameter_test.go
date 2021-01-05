package tree_depth_first_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTreeDiameter(t *testing.T) {
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

 	assert.Equal(t, expected, TreeDiameter(root))
}

func TestTreeDiameter2(t *testing.T) {
	root := &node{
		val:   1,
		left:  &node{
			val:   2,
		},
		right: &node{
			val:   3,
			left:  &node{
				val:   5,
				left: &node{
					val:   7,
				},
				right: &node{
					val:   8,
					left:  &node{
						val:   10,
					},
				},
			},
			right: &node{
				val:   6,
				left: &node{
					val:   9,
					left:  &node{
						val:   11,
					},
				},
			},
		},
	}

	expected := []int{10, 8, 5, 3, 6, 9, 11}

 	assert.Equal(t, expected, TreeDiameter(root))
}
