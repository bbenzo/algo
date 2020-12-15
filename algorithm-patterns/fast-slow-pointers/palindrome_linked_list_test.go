package fast_slow_pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindromeLinkedListTrue(t *testing.T) {
	head := &Node{
		val:  2,
		next: &Node{
			val:  4,
			next: &Node{
				val:  6,
				next: &Node{
					val:  4,
					next: &Node{
						val:  2,
						next: nil,
					},
				},
			},
		},
	}

	assert.True(t, PalindromeLinkedList(head))
}

func TestPalindromeLinkedListFalse(t *testing.T) {
	head := &Node{
		val:  2,
		next: &Node{
			val:  4,
			next: &Node{
				val:  6,
				next: &Node{
					val:  4,
					next: &Node{
						val:  2,
						next: &Node{
							val:  2,
							next: nil,
						},
					},
				},
			},
		},
	}

	assert.False(t, PalindromeLinkedList(head))
}
