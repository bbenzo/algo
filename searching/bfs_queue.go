package searching

import "github.com/pkg/errors"

type queue []*Node

type Queue interface {
	Enqueue(item *Node)
	Dequeue() (*Node, error)
	IsEmpty() bool
}

func New() Queue {
	return &queue{}
}

func (q *queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *queue) Enqueue(item *Node) {
	*q = append(*q, item)
}

func (q *queue) Dequeue() (*Node, error) {
	if len(*q) == 0 {
		return nil, errors.New("queue is empty")
	}

	var item *Node
	*q, item = (*q)[1:], (*q)[0]
	return item, nil
}

func BreadthFirstSearchQueue(root *Node, value int) *Node {
	q := New()
	q.Enqueue(root)

	for !q.IsEmpty() {
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
