package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestSubStringWithKDistinctCharactersDoesNotExist(t *testing.T) {
	word := "abcd"
	k := 5

	result := LongestSubStringWithKDistinctCharacters(word, k)

	assert.Equal(t, 0, len(result))
	assert.Equal(t, "", result)
}

func TestLongestSubStringWithKDistinctCharactersLonger(t *testing.T) {
	word := "abcdefffghijjk"
	k := 5

	result := LongestSubStringWithKDistinctCharacters(word, k)

	assert.Equal(t, 8, len(result))
	assert.Equal(t, "fffghijj", result)
}

func TestLongestSubStringWithKDistinctCharactersLong(t *testing.T) {
	word := "abcdefffg"
	k := 5

	result := LongestSubStringWithKDistinctCharacters(word, k)

	assert.Equal(t, 7, len(result))
	assert.Equal(t, "bcdefff", result)
}

func TestLongestSubStringWithKDistinctCharacters(t *testing.T) {
	word := "baraaci"
	k := 2

	result := LongestSubStringWithKDistinctCharacters(word, k)

	assert.Equal(t, 4, len(result))
	assert.Equal(t, "araa", result)
}


func TestLongestSubStringWithKDistinctCharacters2(t *testing.T) {
	word := "baraaccci"
	k := 2

	result := LongestSubStringWithKDistinctCharacters(word, k)

	assert.Equal(t, 5, len(result))
	assert.Equal(t, "aaccc", result)
}

func TestLongestSubStringWithKDistinctCharactersOneChar(t *testing.T) {
	word := "araaci"
	k := 1

	result := LongestSubStringWithKDistinctCharacters(word, k)

	assert.Equal(t, 2, len(result))
}
