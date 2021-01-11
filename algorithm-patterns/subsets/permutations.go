package subsets

import "errors"

type arrQueue [][]int

type ArrQueue interface {
	Enqueue(item []int)
	Dequeue() ([]int, error)
	Len() int
}

func NewArrQueue() ArrQueue {
	return &arrQueue{}
}

func (q *arrQueue) Len() int {
	return len(*q)
}

func (q *arrQueue) Enqueue(item []int) {
	*q = append(*q, item)
}

func (q *arrQueue) Dequeue() ([]int, error) {
	if len(*q) == 0 {
		return nil, errors.New("nodeQueue is empty")
	}

	var item []int
	*q, item = (*q)[1:], (*q)[0]
	return item, nil
}

func Permutations(arr []int) [][]int {
	result := [][]int{}
	q := NewArrQueue()
	q.Enqueue([]int{})

	for i := range arr {

		length := q.Len()

		for j := 0; j < length; j++ {
			subset, err := q.Dequeue()
			if err != nil {
				panic(err)
			}

			subset = append(subset, arr[i])
			if i < len(arr)-1 {
				q.Enqueue(append([]int{}, subset...))
			} else {
				result = append(result, append([]int{}, subset...))
			}

			k := len(subset) - 1
			for k > 0 {
				subset[k], subset[k-1] = subset[k-1], subset[k]
				if i < len(arr)-1 {
					q.Enqueue(append([]int{}, subset...))
				} else {
					result = append(result, append([]int{}, subset...))
				}
				k--
			}

		}
	}

	return result
}

func factorial(n int) (result int) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}
