package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSmallestStringContainingPattern(t *testing.T) {
	str := "aabdec"
	pattern := "abc"

	expected := "abdec"

	result := SmallestStringContainingPattern(str, pattern)

	assert.Equal(t, expected, result)
}

func TestSmallestStringContainingPattern2(t *testing.T) {
	str := "abdbca"
	pattern :=  "abc"

	expected := "bca"

	result := SmallestStringContainingPattern(str, pattern)

	assert.Equal(t, expected, result)
}

func TestSmallestStringContainingPattern3(t *testing.T) {
	str := ""
	pattern :=  ""

	expected := ""

	result := SmallestStringContainingPattern(str, pattern)

	assert.Equal(t, expected, result)
}
