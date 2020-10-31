package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestSubStringNoRepeat(t *testing.T) {
	input := "aabccbb"
	expected := "abc"

	result := LongestSubStringNoRepeat(input)
	assert.Equal(t, expected, result)
}

func TestLongestSubStringNoRepeat2(t *testing.T) {
	input := "abbbb"
	expected := "ab"

	result := LongestSubStringNoRepeat(input)
	assert.Equal(t, expected, result)
}

func TestLongestSubStringNoRepeat3(t *testing.T) {
	input := "abccde"
	expected := "abc"

	result := LongestSubStringNoRepeat(input)
	assert.Equal(t, expected, result)
}

func TestLongestSubStringNoRepeat4(t *testing.T) {
	input := "abccdeffg"
	expected := "cdef"

	result := LongestSubStringNoRepeat(input)
	assert.Equal(t, expected, result)
}

func TestLongestSubStringNoRepeatEmpty(t *testing.T) {
	input := "aaaa"
	expected := "a"

	result := LongestSubStringNoRepeat(input)
	assert.Equal(t, expected, result)
}
