package two_pointers

func SortedSquares(input []int) []int {
	result := make([]int, len(input))

	start, end := 0, len(input) - 1

	resultIndex := len(input) - 1

	for start < end {
		if uint(input[end]) > uint(start) {
			result[resultIndex] = input[end] * input[end]
			end--
		} else {
			result[resultIndex] = input[start] * input[start]
			start++
		}
		resultIndex--
	}
	return result
}