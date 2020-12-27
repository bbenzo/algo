package cyclic_sort

func FindMissingNum(arr []int) int {
	i := 0
	for i < len(arr) {
		if arr[i] != i && arr[i] < len(arr) {
			arr[i], arr[arr[i]] = arr[arr[i]], arr[i]
		} else {
			i++
		}
	}

	for i := range arr {
		if arr[i] != i {
			return i
		}
	}

	return -1
}
