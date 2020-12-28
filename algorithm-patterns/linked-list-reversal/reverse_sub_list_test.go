package linked_list_reversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseSubLinkedList(t *testing.T) {
	head := &node{
		val:  1,
		next: &node{
			val:  2,
			next: &node{
				val:  3,
				next: &node{
					val:  4,
					next: &node{
						val:  5,
						next: nil,
					},
				},
			},
		},
	}

	expected := &node{
		val:  1,
		next: &node{
			val:  4,
			next: &node{
				val:  3,
				next: &node{
					val:  2,
					next: &node{
						val:  5,
						next: nil,
					},
				},
			},
		},
	}

	reversed := ReverseSubList(head, 2, 4)

	assert.Equal(t, expected, reversed)
}
