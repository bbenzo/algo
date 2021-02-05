package modified_binary_search

func NumberRange(arr []int, key int) []int {
	numRange := []int{-1, -1}

	left, right := 0, len(arr) - 1

	for left <= right {
		pivot := (left + right) / 2

		if arr[pivot] == key {
			first, last := pivot, pivot
			for first > -1 && arr[first] == key {
					first--
			}

			for last < len(arr) && arr[last] == key {
					last++
			}

			numRange[0] = first+1
			numRange[1] = last-1

			return numRange
		}

		if key < arr[pivot] {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

	return numRange
}
