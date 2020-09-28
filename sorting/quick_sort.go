package sorting

import "math/rand"

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// find random pivot
	pivot := rand.Int() % len(arr)

	// move pivot to end of array
	lastIndex := len(arr) - 1
	arr[pivot], arr[lastIndex] = arr[lastIndex], arr[pivot]

	higherThanIndex, unknownIndex := 0, 0

	for unknownIndex < lastIndex {
		if arr[unknownIndex] <= arr[lastIndex] {
			arr[higherThanIndex], arr[unknownIndex] = arr[unknownIndex], arr[higherThanIndex]
			higherThanIndex++
		}
		unknownIndex++
	}

	left, right := QuickSort(arr[:higherThanIndex]), QuickSort(arr[higherThanIndex:])
	return append(left, right...)
}
