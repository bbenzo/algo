package sorting

func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	left, right := divide(arr)

	return merge(MergeSort(left), MergeSort(right))
}

func divide(arr []int) ([]int, []int) {
	length := len(arr)
	m := int(length / 2)
	left := arr[0:m]
	right := arr[m:]
	return left, right
}

func merge(left []int, right []int) []int {
	merged := make([]int, len(left) + len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			merged[i] = left[0]
			left = left[1:]
		} else {
			merged[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
        merged[i] = left[j]
        i++
    }
    for j := 0; j < len(right); j++ {
        merged[i] = right[j]
        i++
    }

	return merged
}
