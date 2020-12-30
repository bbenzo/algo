package breadth_first_search

func BinaryTreeLevelOrderTraverseAverage(root *node) []int {
	var result []int
	q := queue{}
	q.Enqueue(root)

	for q.Len() > 0 {
		length := q.Len()
		sum := 0

		for i := 0; i < length; i++ {
			item, err := q.Dequeue()
			if err != nil {
				panic(err)
			}

			sum += item.val

			if item.left != nil {
				q.Enqueue(item.left)
			}

			if item.right != nil {
				q.Enqueue(item.right)
			}
		}

		result = append(result, sum / length)
	}

	return result
}
