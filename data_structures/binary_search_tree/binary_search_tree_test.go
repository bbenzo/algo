package binary_search_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearchTree_InsertRootNode(t *testing.T) {
	tree := BinarySearchTree{}
	tree.Insert(5)

	assert.Equal(t, 5, tree.Root.Value)
}

func TestBinarySearchTree_InsertSecondSmallerNode(t *testing.T) {
	tree := BinarySearchTree{}
	tree.Insert(5)
	tree.Insert(3)

	assert.Equal(t, 3, tree.Root.Left.Value)
}

func TestBinarySearchTree_InsertSecondLargerNode(t *testing.T) {
	tree := BinarySearchTree{}
	tree.Insert(5)
	tree.Insert(7)

	assert.Equal(t, 7, tree.Root.Right.Value)
}


func TestBinarySearchTree_InsertMultipleNodeSmallerAndLarger(t *testing.T) {
	tree := BinarySearchTree{}
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(8)
	tree.Insert(1)

	assert.Equal(t, 5, tree.Root.Value)
	assert.Equal(t, 4, tree.Root.Left.Value)
	assert.Equal(t, 7, tree.Root.Right.Value)
	assert.Equal(t, 2, tree.Root.Left.Left.Value)
	assert.Equal(t, 6, tree.Root.Right.Left.Value)
	assert.Equal(t, 8, tree.Root.Right.Right.Value)
	assert.Equal(t, 1, tree.Root.Left.Left.Left.Value)
}
