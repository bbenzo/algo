package fast_slow_pointers

func PalindromeLinkedList(head *Node) bool {
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	reversed := reverse(slow)

	for reversed != nil && head != nil {
		if reversed.val != head.val {
			return false
		}

		reversed = reversed.next
		head = head.next
	}

	return true
}

func reverse(head *Node) *Node {
	var prev *Node
	for head != nil {
		next := head.next

		head.next = prev
		prev = head
		head = next
	}

	return prev
}
