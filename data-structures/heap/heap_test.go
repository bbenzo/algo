package heap

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinHeap(t *testing.T) {
	h := &MinIntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)

	assert.Equal(t, (*h)[0], 1)
	assert.Equal(t, (*h)[1], 2)
	assert.Equal(t, (*h)[2], 5)
	assert.Equal(t, (*h)[3], 3)
}

func TestMaxHeap(t *testing.T) {
	h := &MaxIntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)

	assert.Equal(t, (*h)[0], 5)
	assert.Equal(t, (*h)[1], 3)
	assert.Equal(t, (*h)[2], 2)
	assert.Equal(t, (*h)[3], 1)
}
