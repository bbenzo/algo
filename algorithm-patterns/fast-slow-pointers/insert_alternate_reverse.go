package fast_slow_pointers

func InsertAlternateReverse(head *Node) *Node {
	slow, fast := head, head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	middle := slow

	reversed := reverse(middle)

	result := insert(head, reversed)

	return result
}

func insert(head *Node, reversed *Node) *Node {
	result := head
	headToDo := head.next

	result.next = reversed
	reversedToDo := reversed.next

	last := result.next
	for last.next != nil && headToDo.next != nil {
		last.next = headToDo
		headToDo = headToDo.next

		last.next.next = reversedToDo
		reversedToDo = reversedToDo.next

		last = last.next.next

	}

	return result
}
