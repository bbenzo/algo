package bitwise_xor

import "math"

func ComplementBase10(x int) int {
	bitCount := 0
	n := x

	// calculate the max bit by counting
	for n > 0 {
		bitCount++
		n >>= 1
	}


	allBitsSet := int(math.Pow(2, float64(bitCount))) - 1

	return x ^ allBitsSet
}
