package modified_binary_search

func MinimumDifference(arr []int, key int) int {
	left, right := 0, len(arr) - 1

	for left <= right {
		pivot := (left + right) / 2

		if key < arr[pivot] {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

	return arr[right]
}
