package tree_depth_first_search

func TreeDiameter(root *node) []int {
	result := []int{}

	if root == nil {
		return result
	}

	if root.left != nil && root.right != nil {
		diameter := traverseDiameter(root)

		if len(diameter) > len(result) {
			result = diameter
		}
	}

	left := TreeDiameter(root.left)
	right := TreeDiameter(root.right)

	if len(left) > len(result) {
		result = left
	}

	if len(right) > len(result) {
		result = right
	}

	return result
}

func traverseDiameter(root *node) []int {
	diameter := []int{}

	if root == nil {
		return diameter
	}

	left := longestPath(root.left, []int{})

	length := len(left)
	for length > 0 {
		diameter = append(diameter, left[length-1])
		length--
	}

	diameter = append(diameter, root.val)

	right := longestPath(root.right, []int{})

	diameter = append(diameter, right...)

	return diameter
}

func longestPath(root *node, arr []int) []int {
	if root == nil {
		return arr
	}

	arr = append(arr, root.val)

	leftPath := longestPath(root.left, arr)

	rightPath := longestPath(root.right, arr)

	if len(leftPath) > len(rightPath) {
		return leftPath
	}

	return rightPath
}
