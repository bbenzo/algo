package sorting

import "math/rand"

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := divideQuickSort(arr)
	left, right = QuickSort(left), QuickSort(right)
	return append(left, right...)
}

func divideQuickSort(arr []int) ([]int, []int) {
	right := len(arr) - 1
	pivot := rand.Int() % len(arr)

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] > arr[right] {
			arr[i], arr[right] = arr[right], arr[i]
		}
	}

	return arr[:pivot], arr[pivot:]
}
