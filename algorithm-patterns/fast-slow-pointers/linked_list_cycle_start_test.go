package fast_slow_pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedListCycleStart(t *testing.T) {
	node6 := &Node{val: 6, next: nil}
	node5 := &Node{val: 5, next: node6}
	node4 := &Node{val: 4, next: node5}
	node3 := &Node{val: 3, next: node4}
	node2 := &Node{val: 2, next: node3}
	node1 := &Node{val: 1, next: node2}

	node6.next = node3

	assert.Equal(t, node3, LinkedListCycleStart(node1))
}

func TestLinkedListCycleStart2(t *testing.T) {
	node6 := &Node{val: 6, next: nil}
	node5 := &Node{val: 5, next: node6}
	node4 := &Node{val: 4, next: node5}
	node3 := &Node{val: 3, next: node4}
	node2 := &Node{val: 2, next: node3}
	node1 := &Node{val: 1, next: node2}

	node6.next = node4

	assert.Equal(t, node4, LinkedListCycleStart(node1))
}
