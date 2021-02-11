package bitwise_xor

func FlipInvertImage(img [][]int) [][]int {
	for _, line := range img {
		for i, j := 0, len(line) -1; i < j; i, j = i+1, j-1 {
			line[i], line[j] = line[j], line[i]
		}

		for i := range line {
			line[i] ^= 1
		}
	}

	return img
}
