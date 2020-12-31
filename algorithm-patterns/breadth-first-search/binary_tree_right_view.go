package breadth_first_search

func BinaryTreeRightView(root *node) []int {
	var result []int

	q := queue{}
	q.Enqueue(root)

	for q.Len() > 0 {
		item, err := q.Dequeue()
		if err != nil {
			panic(err)
		}

		result = append(result, item.val)

		if item.right != nil {
			q.Enqueue(item.right)
		}
	}

	return result
}
