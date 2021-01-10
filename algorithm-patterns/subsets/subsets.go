package subsets

import (
	"errors"
)

type queue []int

type Queue interface {
	Enqueue(item int)
	Dequeue() (int, error)
	Len() int
}

func NewQueue() Queue {
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
		return 0, errors.New("nodeQueue is empty")
	}

	var item int
	*q, item = (*q)[1:], (*q)[0]
	return item, nil
}

func Subsets(arr []int) [][]int {
	if len(arr) == 0 {
		return nil
	}

	result := [][]int{}
	result = append(result, []int{})


	for _, num := range arr {

		length := len(result)
		for i := 0 ; i < length; i++ {
			newSubset := append([]int{}, result[i]...)
			newSubset = append(newSubset, num)
			result = append(result, newSubset)
		}
	}

	return result
}
