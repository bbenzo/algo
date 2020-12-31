package breadth_first_search

func BinaryTreeLevelOrderSuccessor(root *node, k int) int {
	q := queue{}
	q.Enqueue(root)

	for q.Len() > 0 {
		item, err := q.Dequeue()
		if err != nil {
			panic(err)
		}

		if item.left != nil {
			q.Enqueue(item.left)
		}

		if item.right != nil {
			q.Enqueue(item.right)
		}

		if item.val == k {
			break
		}
	}

	item, err := q.Dequeue()
	if err != nil {
		panic(err)
	}

	return item.val
}
