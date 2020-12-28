package linked_list_reversal

func RotateLinkedList(head *node, k int) *node {
	current := head
	var previous *node
	var next *node

	for k > 0 {
		next = current.next
		previous = current
		previous.next = nil

		current = next
		current.Push(previous)

		k--
	}

	return current
}
