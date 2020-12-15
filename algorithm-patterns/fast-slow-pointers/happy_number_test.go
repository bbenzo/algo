package fast_slow_pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHappyNumber(t *testing.T) {
	assert.True(t, HappyNumber(&Node{val: 23}))
}

func TestHappyNumberNot(t *testing.T) {
	assert.False(t, HappyNumber(&Node{val: 12}))
}