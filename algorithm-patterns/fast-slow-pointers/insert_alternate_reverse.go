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
	nextReverse := reversed
	for head != nil && head.next != nil {
		result.next = nextReverse
	}

	return prev
}
