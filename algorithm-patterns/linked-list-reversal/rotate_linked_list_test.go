package linked_list_reversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotateLinkedList(t *testing.T) {
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
						next: &node{
							val:  6,
							next: nil,
						},
					},
				},
			},
		},
	}

	k := 3

	expected := &node{
		val:  4,
		next: &node{
			val:  5,
			next: &node{
				val:  6,
				next: &node{
					val:  1,
					next: &node{
						val:  2,
						next: &node{
							val:  3,
							next: nil,
						},
					},
				},
			},
		},
	}

	assert.Equal(t, expected, RotateLinkedList(head, k))
}
