package recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterativeFactorial(t *testing.T) {
	result := IterativeFactorial(5)

	assert.Equal(t, 120, result)
}

func TestRecursiveFactorial(t *testing.T) {
	result := RecursiveFactorial(5)

	assert.Equal(t, 120, result)
}

func TestIsPalindrome(t *testing.T) {
	result := IsPalindrome("otto")

	assert.Equal(t, true, result)
}

func TestIsPalindrome2(t *testing.T) {
	result := IsPalindrome("martin nitram")

	assert.Equal(t, true, result)
}

func TestIsPalindromeNot(t *testing.T) {
	result := IsPalindrome("otti")

	assert.Equal(t, false, result)
}

func TestRecursivePow(t *testing.T) {
	result := RecursivePow(2, 4)

	assert.Equal(t, 16, result)
}
