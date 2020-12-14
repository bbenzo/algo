package fast_slow_pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHappyNumber(t *testing.T) {
	assert.True(t, HappyNumber("23", []int{}))
}

func TestHappyNumberNot(t *testing.T) {
	assert.False(t, HappyNumber("12", []int{}))
}