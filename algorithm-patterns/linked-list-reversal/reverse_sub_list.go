package linked_list_reversal

func ReverseSubList(head *node, pos1, pos2 int) *node {
	current := head

	var result *node
	var rest *node

	count := 1
	for count < pos1 && current != nil {
		rest = current.next
		current.next = nil
		if result == nil {
			result = current
		} else {
			result.Push(current)
		}
		count++
	}

	var next *node
	var previous *node
	i := 0
	for rest != nil && i < pos2-pos1+1 {
		next = rest.next
		rest.next = previous
		previous = rest
		rest = next
		i++
	}

	if result == nil {
		result = previous
	} else {
		result.Push(previous)
	}

	result.Push(rest)

	return result
}
