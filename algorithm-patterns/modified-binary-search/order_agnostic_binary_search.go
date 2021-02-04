package modified_binary_search

func OrderAgnosticBinarySearch(arr []int, key int) int {
	isAcsending := true
	if arr[0] > arr[len(arr)-1] {
		isAcsending = false
	}

	left, right := 0, len(arr)-1

	for left <= right {
		pivot := (left + right) / 2

		if arr[pivot] == key {
			return pivot
		}

		if isAcsending {
			if key < arr[pivot] {
				right = pivot - 1
			} else {
				left = pivot + 1
			}
		} else {
			if key > arr[pivot] {
				right = pivot - 1
			} else {
				left = pivot + 1
			}
		}
	}

	return -1
}
