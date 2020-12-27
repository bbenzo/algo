package linked_list_reversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	head := &node{
		val:  1,
		next: &node{
			val:  2,
			next: &node{
				val:  3,
				next: nil,
			},
		},
	}

	expected := &node{
		val:  3,
		next: &node{
			val:  2,
			next: &node{
				val:  1,
				next: nil,
			},
		},
	}

	reversed := Reverse(head)

	assert.Equal(t, expected, reversed)
}
