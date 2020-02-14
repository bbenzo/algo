package sorting

import "math/rand"

func QuickSortRandom(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}

	left, right := divideByPivot(arr, n)

	return merge(QuickSortRandom(left), QuickSortRandom(right))
}

func divideByPivot(arr []int, n int) ([]int, []int) {
	pivot := rand.Int() % n
	return arr[0:pivot], arr[pivot:]
}
