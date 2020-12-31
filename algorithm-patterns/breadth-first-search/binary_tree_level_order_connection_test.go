package breadth_first_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreeLevelOrderConnection(t *testing.T) {
	root := &node{val: 12}
	root.left = &node{val: 7}
	root.right = &node{val: 1}
	root.left.left = &node{val: 9}
	root.right.left = &node{val: 10}
	root.right.right = &node{val: 5}

	expected := []*node{
		&node{val: 12},
		&node{val: 7, left: &node{val: 1}},
		&node{val: 9, left: &node{val: 10, left: &node{val: 5}}},
	}

	assert.Equal(t, expected, BinaryTreeLevelOrderConnection(root))
}
