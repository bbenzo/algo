package linked_list_reversal

type node struct {
	val  int
	next *node
}

func (n *node) Tail() *node {
	for n.next != nil {
		*n = *n.next
	}

	return n
}

func (n *node) Push(newNode *node) {
	for n.next != nil {
		n = n.next
	}

	n.next = newNode

}

func Reverse(current *node) *node {
	var previous *node
	var next *node
	for current != nil {
		next = current.next
		current.next = previous
		previous = current
		current = next
	}

	return previous
}
