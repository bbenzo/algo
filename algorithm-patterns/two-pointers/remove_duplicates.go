package two_pointers

func RemoveDuplicates(input []int) []int {
	i := 0
	nextNonDuplicate := 1

	for i < len(input) {
		if input[nextNonDuplicate - 1] != input[i] {
			input[nextNonDuplicate] = input[i]
			nextNonDuplicate++
		}
		i++
	}

	return input[:nextNonDuplicate]
}
