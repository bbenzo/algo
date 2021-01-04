package tree_depth_first_search

func AllPathsForSum(root *node, k int) [][]int {
	result := [][]int{}
	current := []int{}

	sum := 0
	traverseWithResult(root, k, sum, current, &result)

	return result
}

func traverseWithResult(root *node, k, sum int, current []int, result *[][]int) {
	if root == nil {
		return
	}

	current = append(current, root.val)

	sum += root.val
	if sum == k {
		*result = append(*result, current)
		return
	}

	traverseWithResult(root.left, k, sum, current, result)
	traverseWithResult(root.right, k, sum, current, result)
}
