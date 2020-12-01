package two_pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates1(t *testing.T) {
	input := []int{1, 2, 2, 4}
	expected := []int{1, 2, 4}
	output := RemoveDuplicates(input)

	assert.Equal(t, expected, output)
}

func TestRemoveDuplicates2(t *testing.T) {
	input := []int{1, 1, 2, 4, 6, 6, 9}
	expected := []int{1, 2, 4, 6, 9}
	output := RemoveDuplicates(input)

	assert.Equal(t, expected, output)
}

func TestRemoveDuplicates3(t *testing.T) {
	input := []int{1, 1, 2, 4, 6, 6, 9, 9, 9, 9}
	expected := []int{1, 2, 4, 6, 9}
	output := RemoveDuplicates(input)

	assert.Equal(t, expected, output)
}
