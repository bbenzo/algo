package cyclic_sort

func SmallestMissingPositiveNum(arr []int) int {
	i := 0
	for i < len(arr) {
		if arr[i] - 1 != i && arr[i] - 1 < len(arr) && arr[i] - 1 >= 0 {
			arr[i], arr[arr[i] - 1] = arr[arr[i] - 1], arr[i]
		} else {
			i++
		}
	}

	for i := range arr {
		if arr[i] -1 != i {
			return i+1
		}
	}

	return -1
}
