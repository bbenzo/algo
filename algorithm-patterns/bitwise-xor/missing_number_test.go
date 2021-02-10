package bitwise_xor

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	arr := []int{1, 5, 2, 6, 4}

	assert.Equal(t, 3, MissingNumber(arr))
}

func TestTest(t *testing.T) {
	result := (6 ^ 1) + (1 ^ 6)

	fmt.Println(result)
}
