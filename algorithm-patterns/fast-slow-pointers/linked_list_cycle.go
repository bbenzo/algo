package fast_slow_pointers

type Node struct {
	val int
	next *Node
}

func LinkedListHasCycle(head *Node) bool {
	fast, slow := head, head

	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next

		if slow == fast {
			return true
		}
	}

	return false
}
