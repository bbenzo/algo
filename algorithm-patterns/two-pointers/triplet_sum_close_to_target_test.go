package two_pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTripletSumCloseToTarget(t *testing.T) {
	arr := []int{-2, 0, 1, 2}
	target := 2

	result := TripletSumCloseToTarget(arr, target)

	assert.Equal(t, 1, result)
}

func TestTripletSumCloseToTarget2(t *testing.T) {
	arr := []int{-3, -1, 1, 2}
	target := 1

	result := TripletSumCloseToTarget(arr, target)

	assert.Equal(t, 0, result)
}

func TestTripletSumCloseToTarget3(t *testing.T) {
	arr := []int{1, 0, 1, 1}
	target := 100

	result := TripletSumCloseToTarget(arr, target)

	assert.Equal(t, 3, result)
}
