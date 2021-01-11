package subsets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringCasePermutations(t *testing.T) {
	input := "ad52"
	expected := []string{"ad52", "aD52", "Ad52", "AD52"}

	assert.Equal(t, expected, StringCasePermutations(input))
}

func TestStringCasePermutations2(t *testing.T) {
	input := "ab7c"
	expected := []string{"ab7C", "ab7c", "aB7C", "aB7c", "Ab7C", "Ab7c", "AB7C", "AB7c"}

	assert.Equal(t, expected, StringCasePermutations(input))
}
