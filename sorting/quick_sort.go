package sorting

import "math/rand"

func QuickSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}

	pivot := rand.Int() % n
	left, right := arr[0:pivot], arr[pivot:]

	return merge(QuickSort(left), QuickSort(right))
}
