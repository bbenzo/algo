package fast_slow_pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedListMiddle(t *testing.T) {
	head := &Node{
		val:  1,
		next: &Node{
			val:  2,
			next: &Node{
				val: 3,
				next: &Node{
					val:  4,
					next: &Node{
						val:  5,
					},
				},
			},
		},
	}

	assert.Equal(t, 3, LinkedListMiddle(head))
}

func TestLinkedListMiddle2(t *testing.T) {
	head := &Node{
		val:  1,
		next: &Node{
			val:  2,
			next: &Node{
				val: 3,
				next: &Node{
					val:  4,
					next: &Node{
						val:  5,
						next: &Node{
							val:  6,
							next: nil,
						},
					},
				},
			},
		},
	}

	assert.Equal(t, 4, LinkedListMiddle(head))
}

func TestLinkedListMiddle3(t *testing.T) {
	head := &Node{
		val:  1,
		next: &Node{
			val:  2,
			next: &Node{
				val: 3,
				next: &Node{
					val:  4,
					next: &Node{
						val:  5,
						next: &Node{
							val:  6,
							next: &Node{
								val:  7,
							},
						},
					},
				},
			},
		},
	}

	assert.Equal(t, 4, LinkedListMiddle(head))
}
