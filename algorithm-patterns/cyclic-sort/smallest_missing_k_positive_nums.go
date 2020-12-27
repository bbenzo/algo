package cyclic_sort

func SmallestMissingKPositiveNums(arr []int, k int) []int {
	var result []int

	high := -1
	i := 0
	for i < len(arr) {
		if arr[i] > high {
			high = arr[i]
		}

		if arr[i] - 1 != i && arr[i] - 1 < len(arr) && arr[i] - 1 >= 0 {
			if arr[arr[i] - 1] == arr[i] {
				i++
				continue
			}

			arr[i], arr[arr[i] - 1] = arr[arr[i] - 1], arr[i]
		} else {
			i++
		}
	}

	i = 0
	for i < len(arr) && k > 0 {
		if arr[i] - 1 != i {
			result = append(result, i+1)
			k--
		}
		i++
	}

	i = 0
	for i < k {
		result = append(result, high + i+1)
		i++
	}

	return result
}
