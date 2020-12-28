package linked_list_reversal

func ReverseKSizedSubList(head *node, k int) *node {
	current := head
	var result *node
	var previous *node
	var next *node

	i := 0
	for current != nil {
		next = current.next
		current.next = previous
		previous = current
		current = next
		i++

		if i % k == 0 {
			if result == nil {
				result = previous
			} else {
				result.Push(previous)
			}
			previous = nil
		}
	}

	return result
}
