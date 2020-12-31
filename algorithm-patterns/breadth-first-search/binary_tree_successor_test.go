package breadth_first_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreeSuccessor(t *testing.T) {
	root := &node{val: 12}
	root.left = &node{val: 7}
	root.right = &node{val: 1}
	root.left.left = &node{val: 9}
	root.right.left = &node{val: 10}
	root.right.right = &node{val: 5}

	expected := &node{
		val:   12,
		left:  &node{
			val:   7,
			left:  &node{
				val:   1,
				left:  &node{
					val:   9,
					left:  &node{
						val:   10,
						left:  &node{
							val:   5,
						},
					},
				},
			},
		},
	}

	BinaryTreeSuccessor(root)

	assert.Equal(t, expected, root)
}
