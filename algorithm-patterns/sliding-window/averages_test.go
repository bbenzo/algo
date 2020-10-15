package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAveragesOfSubArrays(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9}

	result := FindAveragesOfSubArrays(arr, 3)
	assert.Equal(t, 3, len(result))
	assert.Equal(t, 3.0, result[0])
	assert.Equal(t, 5.0, result[1])
	assert.Equal(t, 7.0, result[2])
}
