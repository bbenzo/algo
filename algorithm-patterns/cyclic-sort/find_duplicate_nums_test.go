package cyclic_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDuplicateNum(t *testing.T) {
	arr := []int{1, 4, 4, 3, 2}

	assert.Equal(t, 4, FindDuplicateNum(arr))
}

func TestFindDuplicateNum2(t *testing.T) {
	arr := []int{2, 1, 3, 3, 5, 4}

	assert.Equal(t, 3, FindDuplicateNum(arr))
}

func TestFindDuplicateNum3(t *testing.T) {
	arr := []int{2, 4, 1, 4, 4}

	assert.Equal(t, 4, FindDuplicateNum(arr))
}
