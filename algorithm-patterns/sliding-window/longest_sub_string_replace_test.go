package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestSubStringReplace(t *testing.T) {
	input := "abbcb"
	k := 1

	expected := "bbcb"
	result := LongestSubStringReplace(input, k)

	assert.Equal(t, expected, result)
}

func TestLongestSubStringReplace2(t *testing.T) {
	input := "abbcbcccb"
	k := 2

	expected := "cbcccb"
	result := LongestSubStringReplace(input, k)

	assert.Equal(t, expected, result)
}

func TestLongestSubStringReplace3(t *testing.T) {
	input := "abbcbcccbcec"
	k := 3

	expected := "cbcccbcec"
	result := LongestSubStringReplace(input, k)

	assert.Equal(t, expected, result)
}

func TestLongestSubStringReplace4(t *testing.T) {
	input := "aaaa"
	k := 3

	expected := "aaaa"
	result := LongestSubStringReplace(input, k)

	assert.Equal(t, expected, result)
}
