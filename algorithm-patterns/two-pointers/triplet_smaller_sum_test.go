package two_pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTripletsSmallerSum1(t *testing.T) {
	input := []int{-1, 0, 2, 3}
	target  := 3
	expected := 2

	assert.Equal(t, expected, TripletsSmallerSum(input, target))
}

func TestTripletsSmallerSum2(t *testing.T) {
	input := []int{-1, 4, 2, 1, 3}
	target  := 5
	expected := 4

	assert.Equal(t, expected, TripletsSmallerSum(input, target))
}
