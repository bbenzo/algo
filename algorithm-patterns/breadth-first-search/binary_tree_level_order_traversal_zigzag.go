package breadth_first_search

func BinaryTreeLevelOrderTraverseZigzag(root *node) [][]int {
	result := [][]int{}

	q := queue{}
	q.Enqueue(root)

	alternate := false
	for q.Len() > 0 {
		length := q.Len()
		items := make([]int, length)

		i := 0
		for i < length {
			item, err := q.Dequeue()
			if err != nil {
				panic(err)
			}

			if alternate {
				items[length - 1 - i] = item.val
			} else {
				items[i] = item.val
			}

			if item.left != nil {
				q.Enqueue(item.left)
			}

			if item.right != nil {
				q.Enqueue(item.right)
			}

			i++
		}

		result = append(result, items)

		alternate = !alternate
	}

	return result
}
