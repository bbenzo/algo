package modified_binary_search

func CeilingOfANumber(arr []int, key int) int {
	if key > arr[len(arr) - 1] {
		return -1
	}

	left, right := 0, len(arr) - 1

	for left <= right {
		pivot := (left + right) / 2

		if key < arr[pivot] {
			right = pivot - 1
		} else if key > arr[pivot] {
			left = pivot + 1
		} else {
			return pivot
		}
	}

	return left
}
