package sorting

func InsertionSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		for j := i; j > -1; j-- {
			if j > 0 && arr[j-1] > key {
				arr[j] = arr[j - 1]
			} else {
				arr[j] = key
				break
			}
		}
	}

	return arr
}

func Insert(arr []int, key int) []int {
	arr = append(arr, key)

	if len(arr) == 0 {
		return arr
	}

	for j := len(arr) - 1; j > -1; j-- {
		if j > 0 && arr[j-1] > key {
			arr[j] = arr[j-1]
		} else {
			arr[j] = key
			break
		}
	}

	return arr
}
