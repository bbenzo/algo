package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := Queue{}

	q.enqueue(1)

	assert.Equal(t, 1, len(q))
	assert.Equal(t, 1, q[0])

	q.enqueue(8)
	q.enqueue(4)

	assert.Equal(t, 3, len(q))
	assert.Equal(t, 8, q[1])
	assert.Equal(t, 4, q[2])
}

func TestDequeue(t *testing.T) {
	q := Queue{}

	q.enqueue(2)
	q.enqueue(4)
	q.enqueue(7)

	assert.Equal(t, len(q), 3)

	item, err := q.dequeue()

	assert.Nil(t, err)
	assert.Equal(t, 2, item)

	item, err = q.dequeue()

	assert.Nil(t, err)
	assert.Equal(t, 4, item)

	item, err = q.dequeue()

	assert.Nil(t, err)
	assert.Equal(t, 7, item)

	item, err = q.dequeue()

	assert.NotNil(t, err)
}
