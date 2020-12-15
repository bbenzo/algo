package fast_slow_pointers

import (
	"log"
	"strconv"
)

func HappyNumber(head *Node) bool {
	tail := tail(head)

	sum := squaresSum(strconv.Itoa(tail.val))

	if sum == 1 {
		return true
	}

	tail.next = &Node{
		val: sum,
	}

	if hasCycle(head) {
		return false
	}

	return HappyNumber(head)
}

func tail(head *Node) *Node {
	last := head
	for last.next != nil {
		last = last.next
	}
	return last
}

func squaresSum(num string) int {
	result := 0
	for i := range num {
		converted, err := strconv.Atoi(string(num[i]))
		if err != nil {
			log.Fatalf("no number %v", converted)
		}

		result += converted * converted
	}

	return result
}

func hasCycle(head *Node) bool {
	fast, slow := head, head

	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next

		if slow.val == fast.val {
			return true
		}
	}

	return false
}