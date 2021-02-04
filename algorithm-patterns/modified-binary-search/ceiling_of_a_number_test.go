package modified_binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCeilingOfANumber(t *testing.T) {
	input := []int{4, 6, 10}
	key := 6

	result := CeilingOfANumber(input, key)

	assert.Equal(t, 1, result)
}


func TestCeilingOfANumber2(t *testing.T) {
	input := []int{1, 3, 8, 10, 15}
	key := 12

	result := CeilingOfANumber(input, key)

	assert.Equal(t, 4, result)
}

func TestCeilingOfANumber3(t *testing.T) {
	input := []int{4, 6, 10}
	key := 17

	result := CeilingOfANumber(input, key)

	assert.Equal(t, -1, result)
}

func TestCeilingOfANumber4(t *testing.T) {
	input := []int{1, 3, 8, 10}
	key := -1

	result := CeilingOfANumber(input, key)

	assert.Equal(t, 0, result)
}
