package tree_depth_first_search

type node struct {
	val   int
	left  *node
	right *node
}

func BinaryTreePath(root *node, k int) bool {
	return traverse(root, 0, k)
}

func traverse(root *node, sum, k int) bool {
	if root == nil {
		return false
	}

	sum += root.val
	if sum == k {
		return true
	}

	return traverse(root.left, sum, k) || traverse(root.right, sum, k)
}
