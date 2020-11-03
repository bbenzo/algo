package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsPermutationTrue(t *testing.T) {
	str := "oiddbcaf"
	pattern := "abc"

	assert.True(t, ContainsPermutation(str, pattern))
}

func TestContainsPermutationFalse(t *testing.T) {
	str := "odicf"
	pattern := "dc"

	assert.False(t, ContainsPermutation(str, pattern))
}

func TestContainsPermutationTrueSincePatternAndStringArePermutationOfEachOther(t *testing.T) {
	str := "bcdxabcdy"
	pattern := "bcdyabcdx"

	assert.True(t, ContainsPermutation(str, pattern))
}

func TestContainsPermutationTrueOnceAgain(t *testing.T) {
	str := "aaacb"
	pattern := "abc"

	assert.True(t, ContainsPermutation(str, pattern))
}
