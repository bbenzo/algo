package breadth_first_search

func BinaryTreeSuccessor(root *node) {
	q := queue{}
	q.Enqueue(root)

	var current, previous *node
	for q.Len() > 0 {
		var err error
		current, err = q.Dequeue()
		if err != nil {
			panic(err)
		}

		if previous != nil {
			previous.left = current
		}

		previous = current

		if current.left != nil {
			q.Enqueue(current.left)
			current.left = nil
		}

		if current.right != nil {
			q.Enqueue(current.right)
			current.right = nil
		}
	}
}
