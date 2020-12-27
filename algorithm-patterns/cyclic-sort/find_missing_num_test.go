package cyclic_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMissingNum(t *testing.T) {
	arr := []int{4, 0, 3, 1}

	assert.Equal(t, 2, FindMissingNum(arr))
}

func TestFindMissingNum2(t *testing.T) {
	arr := []int{8, 3, 5, 2, 4, 6, 0, 1}

	assert.Equal(t, 7, FindMissingNum(arr))
}
