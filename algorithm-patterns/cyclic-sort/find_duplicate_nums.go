package cyclic_sort

func FindDuplicateNum(arr []int) int {
	i := 0
	for i < len(arr) {
		if arr[i] - 1 != i {
			if arr[i] == arr[arr[i] - 1] {
				return arr[i]
			}

			arr[i], arr[arr[i] - 1] = arr[arr[i] - 1], arr[i]
		} else {
			i++
		}
	}

	return -1
}
