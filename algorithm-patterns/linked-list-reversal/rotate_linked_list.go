package linked_list_reversal

func RotateLinkedList(head *node, k int) *node {
	if head == nil || head.next == nil {
		return head
	}

	listLength := 1
	lastNode := head
	for lastNode.next != nil {
		lastNode = lastNode.next
		listLength++
	}

	lastNode.next = head
	k %= listLength

	for i := 0; i < k; i++ {
		lastNode = lastNode.next
	}

	head = lastNode.next
	lastNode.next = nil

	return head
}
