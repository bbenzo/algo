package dijkstras

import "github.com/pkg/errors"

type queue []*Node

type Queue interface {
	Enqueue(item *Node)
	Dequeue() (*Node, error)
	Len() int
}

func NewQueue() Queue {
	return &queue{}
}

func (q *queue) Len() int {
	return len(*q)
}

func (q *queue) Enqueue(item *Node) {
	*q = append(*q, item)
}

func (q *queue) Dequeue() (*Node, error) {
	if len(*q) == 0 {
		return nil, errors.New("nodeQueue is empty")
	}

	var item *Node
	*q, item = (*q)[1:], (*q)[0]
	return item, nil
}