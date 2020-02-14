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
