package cyclic_sort

func FindMissingNums(arr []int) []int {
	var result []int

	i := 0
	for i < len(arr) {
		if arr[i] - 1 != i && arr[arr[i] - 1] != arr[i] {
			arr[i], arr[arr[i] - 1] = arr[arr[i] - 1], arr[i]
		} else {
			i++
		}
	}

	for i := range arr {
		if arr[i] - 1 != i {
			result = append(result, i+1)
		}
	}

	return result
}
