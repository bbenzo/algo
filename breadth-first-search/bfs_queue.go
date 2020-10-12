package breadth_first_search

import (
	queue2 "github/bbenzo/algo/data_structures/queue"
)

func BreadthFirstSearchQueue(root *queue2.Node, value int) *queue2.Node {
	q := queue2.NewNodeQueue()
	q.Enqueue(root)

	for q.Len() != 0 {
		node, _ := q.Dequeue()

		if node.Index == value {
			return node
		}

		for _, neighbor := range node.Neighbors {
			neighbor.Predecessor = node
			neighbor.Distance = node.Distance + 1
			q.Enqueue(neighbor)
		}
	}

	return nil
}
