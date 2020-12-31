package breadth_first_search

func BinaryTreeLevelOrderConnection(root *node) []*node {
	var result []*node

	q := queue{}
	q.Enqueue(root)

	for q.Len() > 0 {
		length := q.Len()
		var level []*node

		for i := 0; i < length; i++ {
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

			item.left = nil
			item.right = nil
			level = append(level, item)
		}

		for i := 1; i < len(level); i++ {
			level[i-1].left = level[i]
		}

		result = append(result, level[0])
	}

	return result
}
