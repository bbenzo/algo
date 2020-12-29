package breadth_first_search

import "github.com/pkg/errors"

type node struct {
	val   int
	left  *node
	right *node
}

type queue []*node

type Queue interface {
	Enqueue(item *node)
	Dequeue() (*node, error)
	Len() int
}

func NewQueue() Queue {
	return &queue{}
}

func (q *queue) Len() int {
	return len(*q)
}

func (q *queue) Enqueue(item *node) {
	*q = append(*q, item)
}

func (q *queue) Dequeue() (*node, error) {
	if len(*q) == 0 {
		return nil, errors.New("nodeQueue is empty")
	}

	var item *node
	*q, item = (*q)[1:], (*q)[0]
	return item, nil
}

func BinaryTreeLevelOrderTraverse(root *node) [][]int {
	result := [][]int{}

	q := queue{}
	q.Enqueue(root)

	for q.Len() > 0 {
		length := q.Len()
		levelItems := make([]int, length)

		i := 0
		for i < length {
			item, err := q.Dequeue()
			if err != nil {
				panic(err)
			}

			levelItems[i] = item.val

			if item.left != nil {
				q.Enqueue(item.left)
			}

			if item.right != nil {
				q.Enqueue(item.right)
			}

			i++
		}

		result = append(result, levelItems)
	}

	return result
}
