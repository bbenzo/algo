package cyclic_sort

func FindDuplicateNums(arr []int) []int {
	var result []int

	i := 0
	for i < len(arr) {
		if arr[i] - 1 != i {
			if arr[i] == arr[arr[i] - 1] {
				result = append(result, arr[i])
				i++
				continue
			}

			arr[i], arr[arr[i] - 1] = arr[arr[i] - 1], arr[i]
		} else {
			i++
		}
	}

	return result
}
