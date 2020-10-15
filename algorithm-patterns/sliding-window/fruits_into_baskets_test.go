package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFruitsIntoBaskets(t *testing.T) {
	baskets := 2
	given := []string{"A", "B", "C", "A", "C"}
	expected := []string{"C", "A", "C"}

	result := FruitsIntoBaskets(given, baskets)

	assert.Equal(t, len(result), 3)
	assert.Equal(t, result, expected)
}

func TestFruitsIntoBasketsLong(t *testing.T) {
	baskets := 2
	given := []string{"A", "B", "C", "B", "B", "C"}
	expected := []string{"B", "C", "B", "B", "C"}

	result := FruitsIntoBaskets(given, baskets)

	assert.Equal(t, len(result), 5)
	assert.Equal(t, result, expected)
}

func TestFruitsIntoBasketsLonger(t *testing.T) {
	baskets := 5
	given := []string{"A", "B", "C", "B", "B", "C", "D", "E", "E", "F", "F", "H"}
	expected := []string{"B", "C", "B", "B", "C", "D", "E", "E", "F", "F"}

	result := FruitsIntoBaskets(given, baskets)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result, expected)
}
