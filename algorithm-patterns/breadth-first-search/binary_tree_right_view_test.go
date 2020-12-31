package breadth_first_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreeRightView(t *testing.T) {
	root := &node{val: 12}
	root.left = &node{val: 7}
	root.right = &node{val: 1}
	root.left.left = &node{val: 9}
	root.right.left = &node{val: 10}
	root.right.right = &node{val: 5}

	expected := []int{12, 1, 5}

	assert.Equal(t, expected, BinaryTreeRightView(root))
}
