package modified_binary_search

func SearchBitonicArray(arr []int, key int) int {
	left, right := 0, len(arr) - 1

	for left < right {
		mid := (left + right) / 2

		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	max := left
	left, right = 0, max

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == key {
			return mid
		}

		if key < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	left, right = max, len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == key {
			return mid
		}

		if key < arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
