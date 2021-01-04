package tree_depth_first_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumOfAllPathNums(t *testing.T) {
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

	expected := 408

 	assert.Equal(t, expected, SumOfAllPathNums(root))
}
