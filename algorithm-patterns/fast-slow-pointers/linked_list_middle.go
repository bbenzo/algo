package fast_slow_pointers

func LinkedListMiddle(head *Node) int {
	slow, fast := head, head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow.val
}
