package cyclic_sort

func FindCorruptPair(arr []int) []int {
	result := make([]int, 2)

	i := 0
	for i < len(arr) {
		if arr[i] - 1 != i {
			if arr[i] == arr[arr[i] - 1] {
				result[0] = arr[arr[i] - 1]
				result[1] = i+1
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
