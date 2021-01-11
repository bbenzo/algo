package subsets

import (
	"errors"
	"strings"
)

type stringQueue []string

type StringQueue interface {
	Enqueue(item string)
	Dequeue() (string, error)
	Len() int
}

func NewStringQueue() StringQueue {
	return &stringQueue{}
}

func (q *stringQueue) Len() int {
	return len(*q)
}

func (q *stringQueue) Enqueue(item string) {
	*q = append(*q, item)
}

func (q *stringQueue) Dequeue() (string, error) {
	if len(*q) == 0 {
		return "", errors.New("nodeQueue is empty")
	}

	var item string
	*q, item = (*q)[1:], (*q)[0]
	return item, nil
}

func StringCasePermutations(input string) []string {
	q := NewStringQueue()
	q.Enqueue("")

	for i := range input {
		length := q.Len()

		for j := 0; j < length; j++ {
			item, err := q.Dequeue()
			if err != nil {
				panic(err)
			}

			var newItems []string
			if len(item) > 0 && isLowerChar(item[len(item)-1]) {
				newItems = append(newItems, item+string(input[i]))
				newItems = append(newItems, item[:len(item)-1]+strings.ToUpper(string(item[len(item)-1]))+string(input[i]))
			} else {
				newItems = append(newItems, item+string(input[i]))
			}

			for _, newItem := range newItems {
				if i == len(input) - 1 && isLowerChar(input[i]) {
					q.Enqueue(newItem[:len(newItem)-1] +strings.ToUpper(string(newItem[len(newItem)-1])))
				}

				q.Enqueue(newItem)
			}
		}
	}

	result := make([]string, q.Len())
	i := 0
	for q.Len() > 0 {
		item, err := q.Dequeue()
		if err != nil {
			panic(err)
		}

		result[i] = item
		i++
	}

	return result
}

func isLowerChar(input byte) bool {
	if input < 97 || input > 122 {
		return false
	}
	return true
}
