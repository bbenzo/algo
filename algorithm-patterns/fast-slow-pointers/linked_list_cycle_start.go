package fast_slow_pointers

func LinkedListCycleStart(head *Node) *Node {
	fast, slow := head, head

	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next

		if slow == fast {
			return slow
		}

		if fast.next == slow {
			return fast
		}
	}

	return head
}
