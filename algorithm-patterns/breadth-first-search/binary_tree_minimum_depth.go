package breadth_first_search

func BinaryTreeMinimumDepth(root *node) int {
	depth := 0

	q := queue{}
	q.Enqueue(root)

	for q.Len() > 0 {
		depth++

		length := q.Len()

		for i := 0; i < length; i++ {
			item, err :=  q.Dequeue()
			if err != nil {
				panic(err)
			}

			if item.left == nil && item.right == nil {
				return depth
			}

			if item.left != nil {
				q.Enqueue(item.left)
			}

			if item.right != nil {
				q.Enqueue(item.right)
			}
		}
	}

	return depth
}
