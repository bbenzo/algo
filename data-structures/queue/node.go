package queue

import "github.com/pkg/errors"

type Node struct {
	Index             int
	Predecessor       *Node
	Distance          int
	Neighbors         []*Node
}

type nodeQueue []*Node

type NodeQueue interface {
	Enqueue(item *Node)
	Dequeue() (*Node, error)
	Len() int
}

func NewNodeQueue() NodeQueue {
	return &nodeQueue{}
}

func (q *nodeQueue) Len() int {
	return len(*q)
}

func (q *nodeQueue) Enqueue(item *Node) {
	*q = append(*q, item)
}

func (q *nodeQueue) Dequeue() (*Node, error) {
	if len(*q) == 0 {
		return nil, errors.New("nodeQueue is empty")
	}

	var item *Node
	*q, item = (*q)[1:], (*q)[0]
	return item, nil
}
