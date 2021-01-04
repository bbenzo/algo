package tree_depth_first_search

func CountAllPathsForSum(root *node, sum int) int {
	count := 0
	current := 0
	path := []int{}

	traverseCountAllPathsForSum(root, sum, current, &count, path)
	return count
}

func traverseCountAllPathsForSum(root *node, sum, current int, count *int, path []int) {
	if root == nil {
		return
	}

	current += root.val
	path = append(path, root.val)

	if current == sum {
		*count++
	} else if current > sum {
		subSum := current
		i := 0
		for i < len(path) {
			subSum -= path[i]
			if subSum == sum  {
				*count++
			}
			i++
		}
	}

	if root.left != nil {
		traverseCountAllPathsForSum(root.left, sum, current, count, path)
	}

	if root.right != nil {
		traverseCountAllPathsForSum(root.right, sum, current, count, path)
	}
}
