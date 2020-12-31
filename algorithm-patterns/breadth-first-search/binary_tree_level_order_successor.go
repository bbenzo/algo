package breadth_first_search

func BinaryTreeLevelOrderSuccessor(root *node, k int) int {
	q := queue{}
	q.Enqueue(root)

	var last *node
	for q.Len() > 0 {
		length := q.Len()

		i := 0
		for i < length {
			item, err := q.Dequeue()
			if err != nil {
				panic(err)
			}

			if last != nil && last.val == k {
				return item.val
			} else {
				last = item
			}

			if item.left != nil {
				q.Enqueue(item.left)
			}

			if item.right != nil {
				q.Enqueue(item.right)
			}

			i++
		}
	}

	return -1
}
