package modified_binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextLetter(t *testing.T) {
	input := []rune{'a', 'c', 'f', 'h'}
	key := 'f'

	result := NextLetter(input, key)

	assert.Equal(t, 'h', result)
}

func TestNextLetter2(t *testing.T) {
	input := []rune{'a', 'c', 'f', 'h'}
	key := 'b'

	result := NextLetter(input, key)

	assert.Equal(t, 'c', result)
}

func TestNextLetter3(t *testing.T) {
	input := []rune{'a', 'c', 'f', 'h'}
	key := 'm'

	result := NextLetter(input, key)

	assert.Equal(t, 'a', result)
}

func TestNextLetter4(t *testing.T) {
	input := []rune{'a', 'c', 'f', 'h'}
	key := 'h'

	result := NextLetter(input, key)

	assert.Equal(t, 'a', result)
}
