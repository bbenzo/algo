package tree_depth_first_search

func PathMaxSum(root *node) ([]int, int) {
	var result []int
	var maxSum int

	if root == nil {
		return result, maxSum
	}

	diam, sum := maxPath(root)
	if sum > maxSum {
		maxSum = sum
		result = diam
	}

	leftPath, leftSum := PathMaxSum(root.left)
	rightPath, rightSum := PathMaxSum(root.right)

	if leftSum > maxSum {
		maxSum = leftSum
		result = leftPath
	}

	if rightSum > maxSum {
		maxSum = rightSum
		result = rightPath
	}

	return result, maxSum
}

func maxPath(root *node) ([]int, int) {
	var diam []int
	var sum int
	if root == nil {
		return diam, sum
	}

	leftPath, leftSum := calcMaxPath(root.left, []int{}, 0)
	length := len(leftPath)
	for length > 0 {
		diam = append(diam, leftPath[length-1])
		length--
	}
	sum += leftSum

	diam = append(diam, root.val)
	sum += root.val

	rightPath, rightSum := calcMaxPath(root.right, []int{}, 0)
	diam = append(diam, rightPath...)
	sum += rightSum

	return diam, sum
}

func calcMaxPath(root *node, arr []int, sum int) ([]int, int) {
	if root == nil {
		return arr, sum
	}

	arr = append(arr, root.val)
	sum += root.val

	leftPath, leftSum := calcMaxPath(root.left, arr, sum)

	rightPath, rightSum := calcMaxPath(root.right, arr, sum)

	if leftSum > rightSum {
		return leftPath, leftSum
	}

	return rightPath, rightSum
}
