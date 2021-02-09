package modified_binary_search

func RotationCount(arr []int) int {
	start, end := 0, len(arr)-1

	for start <= end {
		mid := (start + end) / 2

		if arr[start] > arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	if start == len(arr) {
		return 0
	}

	return start
}
