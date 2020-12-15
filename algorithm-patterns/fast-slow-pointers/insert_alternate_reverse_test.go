package fast_slow_pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertAlternateReverse(t *testing.T) {
	head := &Node{
		val: 2,
		next: &Node{
			val: 4,
			next: &Node{
				val: 6,
				next: &Node{
					val: 8,
					next: &Node{
						val: 10,
						next: &Node{
							val:  12,
							next: nil,
						},
					},
				},
			},
		},
	}

	expected := &Node{
		val: 2,
		next: &Node{
			val: 12,
			next: &Node{
				val: 4,
				next: &Node{
					val: 10,
					next: &Node{
						val: 6,
						next: &Node{
							val:  8,
							next: nil,
						},
					},
				},
			},
		},
	}

	assert.Equal(t, expected, InsertAlternateReverse(head))
}

func TestInsertAlternateReverse2(t *testing.T) {
	head := &Node{
		val: 2,
		next: &Node{
			val: 4,
			next: &Node{
				val: 6,
				next: &Node{
					val: 8,
					next: &Node{
						val:  10,
						next: nil,
					},
				},
			},
		},
	}

	expected := &Node{
		val: 2,
		next: &Node{
			val: 10,
			next: &Node{
				val: 4,
				next: &Node{
					val: 8,
					next: &Node{
						val:  6,
						next: nil,
					},
				},
			},
		},
	}

	assert.Equal(t, expected, InsertAlternateReverse(head))
}
