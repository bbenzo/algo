package linked_list_reversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseKSizedSubList(t *testing.T) {
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
							next: &node{
								val:  7,
								next: &node{
									val:  8,
									next: &node{
										val:  9,
										next: nil,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	expected := &node{
		val:  3,
		next: &node{
			val:  2,
			next: &node{
				val:  1,
				next: &node{
					val:  6,
					next: &node{
						val:  5,
						next: &node{
							val:  4,
							next: &node{
								val:  9,
								next: &node{
									val:  8,
									next: &node{
										val:  7,
										next: nil,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	reversed := ReverseKSizedSubList(head, 3)

	assert.Equal(t, expected, reversed)
}
