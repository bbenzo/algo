package breadth_first_search

func BinaryTreeLevelOrderTraverseReverse(root *node) [][]int {
	result := [][]int{}

	q := queue{}
	q.Enqueue(root)

	for q.Len() > 0 {
		length := q.Len()
		items := make([]int, length)

		i := 0
		for i < length {
			item, err := q.Dequeue()
			if err != nil {
				panic(err)
			}

			items[i] = item.val

			if item.left != nil {
				q.Enqueue(item.left)
			}

			if item.right != nil {
				q.Enqueue(item.right)
			}

			i++
		}

		result = append([][]int{items}, result...)
	}

	return result
}
