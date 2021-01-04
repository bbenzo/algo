package tree_depth_first_search

import "strconv"

func SumOfAllPathNums(root *node) int {
	sum := 0

	traversePathNums(root, "", &sum)

	return sum
}

func traversePathNums(root *node, current string, sum *int) {
	if root == nil {
		return
	}

	current += strconv.Itoa(root.val)

	if root.left == nil && root.right == nil {
		intVal, _ := strconv.Atoi(current)
		*sum += intVal
		return
	}

	if root.left != nil {
		traversePathNums(root.left, current, sum)
	}

	if root.right != nil {
		traversePathNums(root.right, current, sum)
	}
}
