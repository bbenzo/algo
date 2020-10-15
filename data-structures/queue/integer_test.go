package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := NewIntegerQueue()

	q.Enqueue(1)
	q.Enqueue(8)
	q.Enqueue(4)

	assert.Equal(t, 3, q.Len())
}

func TestDequeue(t *testing.T) {
	q := NewIntegerQueue()

	q.Enqueue(2)
	q.Enqueue(4)
	q.Enqueue(7)

	item, err := q.Dequeue()

	assert.Nil(t, err)
	assert.Equal(t, 2, item)

	item, err = q.Dequeue()

	assert.Nil(t, err)
	assert.Equal(t, 4, item)

	item, err = q.Dequeue()

	assert.Nil(t, err)
	assert.Equal(t, 7, item)

	item, err = q.Dequeue()

	assert.NotNil(t, err)
}
