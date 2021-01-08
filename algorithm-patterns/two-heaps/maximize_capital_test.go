package two_heaps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximizeCapital(t *testing.T) {
	capitals := []int{0, 1, 2}
	profits := []int{1, 2, 3}
	initialCapital := 1
	projectNums := 2

	expected := 6

	assert.Equal(t, expected, MaximizeCapital(capitals, profits, initialCapital, projectNums))
}

func TestMaximizeCapital2(t *testing.T) {
	capitals := []int{0, 1, 2, 3}
	profits := []int{1, 2, 3, 5}
	initialCapital := 0
	projectNums := 3

	expected := 8

	assert.Equal(t, expected, MaximizeCapital(capitals, profits, initialCapital, projectNums))
}
