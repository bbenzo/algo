package modified_binary_search

func SearchRotatedArray(arr []int, key int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == key {
			return mid
		}

		if arr[left] <= arr[mid] {
			if key < arr[mid] && key >= arr[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if key > arr[mid] && key <= arr[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
