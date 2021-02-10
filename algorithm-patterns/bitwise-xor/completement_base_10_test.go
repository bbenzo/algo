package bitwise_xor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComplementBase10_1(t *testing.T) {
	assert.Equal(t, 7, ComplementBase10(8))
}


func TestComplementBase10_2(t *testing.T) {
	assert.Equal(t, 5, ComplementBase10(10))
}
