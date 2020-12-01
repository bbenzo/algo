package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcatenatedWords(t *testing.T) {
	str := "catfoxcat"
	words := []string{"cat", "fox"}

	expected := []int{0, 3}

	result := ConcatenatedWords(str, words)

	assert.Equal(t, expected, result)
}

func TestConcatenatedWords2(t *testing.T) {
	str := "catcatfoxfox"
	words := []string{"cat", "fox"}

	expected := []int{3}

	result := ConcatenatedWords(str, words)

	assert.Equal(t, expected, result)
}
