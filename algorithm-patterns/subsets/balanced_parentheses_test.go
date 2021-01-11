package subsets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBalancedParentheses(t *testing.T) {
	expected := []string{"(())", "()()"}

	assert.Equal(t, expected, BalancedParentheses(2))
}


func TestBalancedParentheses2(t *testing.T) {
	expected := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}

	assert.Equal(t, expected, BalancedParentheses(3))
}
