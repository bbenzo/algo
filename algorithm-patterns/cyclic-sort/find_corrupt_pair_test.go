package cyclic_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindCorruptPair(t *testing.T) {
	arr := []int{3, 1, 2, 5, 2}
	expected := []int{2, 4}

	assert.Equal(t, expected, FindCorruptPair(arr))
}

func TestFindCorruptPair2(t *testing.T) {
	arr := []int{3, 1, 2, 3, 6, 4}
	expected := []int{3, 5}

	assert.Equal(t, expected, FindCorruptPair(arr))
}
