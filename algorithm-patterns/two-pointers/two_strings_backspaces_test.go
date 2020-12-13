package two_pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoStringsBackspaces(t *testing.T) {
	str1 := "xy#z"
	str2 := "xzz#"

	assert.True(t, TwoStringsBackspaces(str1, str2))
}

func TestTwoStringsBackspaces2(t *testing.T) {
	str1 := "xy#z"
	str2 := "xyz#"

	assert.False(t, TwoStringsBackspaces(str1, str2))
}

func TestTwoStringsBackspaces3(t *testing.T) {
	str1 := "xp#"
	str2 := "xyz##"

	assert.True(t, TwoStringsBackspaces(str1, str2))
}

func TestTwoStringsBackspaces4(t *testing.T) {
	str1 := "xywrrmp"
	str2 := "xywrrmu#p"

	assert.True(t, TwoStringsBackspaces(str1, str2))
}
