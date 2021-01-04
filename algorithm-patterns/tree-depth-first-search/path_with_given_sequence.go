package tree_depth_first_search

func PathWithGivenSequence(root *node, seq []int) bool {
	return traversePathWithGivenSequence(root, 0, seq)
}

func traversePathWithGivenSequence(root *node, current int, seq []int) bool {
	if root == nil || root.val != seq[current]{
		return false
	}

	if root.left == nil && root.right == nil {
		return true
	}

	current++

	return traversePathWithGivenSequence(root.left, current, seq) ||
		traversePathWithGivenSequence(root.right, current, seq)
}
