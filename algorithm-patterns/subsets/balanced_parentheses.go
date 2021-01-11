package subsets

import "errors"

const (
	open   = "("
	closed = ")"
)

type Parentheses struct {
	val         string
	openCount   int
	closedCount int
}

type parenthesesQueue []*Parentheses

type ParenthesesQueue interface {
	Enqueue(item *Parentheses)
	Dequeue() (*Parentheses, error)
	Len() int
}

func NewParenthesesQueue() ParenthesesQueue {
	return &parenthesesQueue{}
}

func (q *parenthesesQueue) Len() int {
	return len(*q)
}

func (q *parenthesesQueue) Enqueue(item *Parentheses) {
	*q = append(*q, item)
}

func (q *parenthesesQueue) Dequeue() (*Parentheses, error) {
	if len(*q) == 0 {
		return nil, errors.New("nodeQueue is empty")
	}

	var item *Parentheses
	*q, item = (*q)[1:], (*q)[0]
	return item, nil
}

func BalancedParentheses(n int) []string {
	var result []string

	q := NewParenthesesQueue()
	q.Enqueue(&Parentheses{val: "", openCount: 0, closedCount: 0})

	for q.Len() > 0 {
		length := q.Len()

		for j := 0; j < length; j++ {
			item, err := q.Dequeue()
			if err != nil {
				panic(err)
			}

			if item.openCount < n {
				parentheses := &Parentheses{
					val:         item.val + "(",
					openCount:   item.openCount + 1,
					closedCount: item.closedCount,
				}
				q.Enqueue(parentheses)

				if parentheses.closedCount == n && parentheses.openCount == n {
					result = append(result, parentheses.val)
				}
			}

			if item.openCount > item.closedCount && item.closedCount < n {
				parentheses := &Parentheses{
					val:         item.val + ")",
					openCount:   item.openCount,
					closedCount: item.closedCount + 1,
				}
				q.Enqueue(parentheses)

				if parentheses.closedCount == n && parentheses.openCount == n {
					result = append(result, parentheses.val)
				}
			}
		}
	}

	return result
}
