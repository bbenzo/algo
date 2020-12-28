package linked_list_reversal

func ReverseAlternateKSizedSubList(head *node, k int) *node {
	current := head
	var result *node
	var previous *node
	var next *node
	alternate := true

	count := 0
	for current != nil {
		count++

		if alternate {
			next = current.next
			current.next = previous
			previous = current
			current = next

			if count % k == 0 {
				if result == nil {
					result = previous
				} else {
					result.Push(previous)
				}
				previous = nil
				alternate = !alternate
			}
		} else {
			next = current.next
			current.next = nil
			result.Push(current)
			current = next

			if count % k == 0 {
				alternate = !alternate
			}
		}
	}

	return result
}
