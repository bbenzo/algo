package two_pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniqueQuadruplets(t *testing.T) {
	input := []int{4, 1, 2, -1, 1, -3}
	target := 1

	expected := [][]int{{-3, -1, 1, 4}, {-3, 1, 1, 2}}

	assert.Equal(t, expected, UniqueQuarduplets(input, target))
}
