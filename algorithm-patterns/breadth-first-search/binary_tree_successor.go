package breadth_first_search

func BinaryTreeSuccessor(root *node) *node {
	var result *node

	q := queue{}
	q.Enqueue(root)

	for q.Len() > 0 {
		item, err := q.Dequeue()
		if err != nil {
			panic(err)
		}

		if item.left != nil {
			q.Enqueue(item.left)
			item.left = nil
		}

		if item.right != nil {
			q.Enqueue(item.right)
			item.right = nil
		}

		if result == nil {
			result = item
		} else {
			result.PushLeft(item)
		}
	}

	return result
}
