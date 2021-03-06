package tree_depth_first_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPathWithGivenSequence(t *testing.T) {
	root := &node{
		val:   1,
		left:  &node{
			val:   7,
		},
		right: &node{
			val:   9,
			left:  &node{
				val:   2,
			},
			right: &node{
				val:   9,
			},
		},
	}

	seq := []int{1, 9, 9}

 	assert.True(t, PathWithGivenSequence(root, seq))
}
