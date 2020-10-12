package queue

import "github.com/pkg/errors"

type queue []int

type Queue interface {
	Enqueue(item int)
	Dequeue() (int, error)
	Len() int
}

func New() Queue {
	return &queue{}
}

func (q *queue) Len() int {
	return len(*q)
}

func (q *queue) Enqueue(item int) {
	*q = append(*q, item)
}

func (q *queue) Dequeue() (int, error) {
	if len(*q) == 0 {
		return 0, errors.New("queue is empty")
	}

	var item int
	*q, item = (*q)[1:], (*q)[0]
	return item, nil
}
