package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringAnagramIndexes(t *testing.T) {
	str := "ppqp"
	pattern := "pq"
	expected := []int{1, 2}
	result := StringAnagramIndexes(str, pattern)

	assert.Equal(t, expected, result)
}

func TestStringAnagramIndexes2(t *testing.T) {
	str := "abbcabc"
	pattern := "abc"
	expected := []int{2, 3, 4}
	result := StringAnagramIndexes(str, pattern)

	assert.Equal(t, expected, result)
}
