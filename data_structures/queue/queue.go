package queue

import "github.com/pkg/errors"

type Queue []int

func (q *Queue) enqueue(item int) {
	*q = append(*q, item)
}

func (q *Queue) dequeue() (int, error) {
	if len(*q) == 0 {
		return 0, errors.New("queue is empty")
	}

	var item int
	*q, item = (*q)[1:], (*q)[0]
	return item, nil
}
