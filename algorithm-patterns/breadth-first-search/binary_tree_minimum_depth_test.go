package breadth_first_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreeMinimumDepth(t *testing.T) {
	root := &node{val: 12}
	root.left = &node{val: 7}
	root.right = &node{val: 1}
	root.right.left = &node{val: 10}
	root.right.right = &node{val: 5}

	expected := 2

	assert.Equal(t, expected, BinaryTreeMinimumDepth(root))
}
