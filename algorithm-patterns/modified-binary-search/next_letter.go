package modified_binary_search

func NextLetter(arr []rune, key rune) rune {
	left, right := 0, len(arr)-1

	for left <= right {
		pivot := (left + right) / 2

		if key == arr[pivot] {
			if pivot == len(arr) -1 {
				return arr[0]
			}

			return arr[pivot + 1]
		}

		if key < arr[pivot] {
			right = pivot - 1
		} else if key > arr[pivot] {
			left = pivot + 1
		}
	}

	if left >= len(arr) {
		return arr[0]
	}

	return arr[left]
}
