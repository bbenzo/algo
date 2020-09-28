package sorting

import "math/rand"

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// find random pivot
	pivot := rand.Int() % len(arr)

	// move pivot to end of array
	high := len(arr) - 1
	arr[pivot], arr[high] = arr[high], arr[pivot]

	// define beginning index of array higher than pivot and beginning index of array which has not been compared to pivot yet
	higherThanIndex, unknownIndex := 0, 0

	for unknownIndex < high {
		// if element is smaller than pivot, swap element with first element of array higher(!) than pivot and move index of that one up
		if arr[unknownIndex] <= arr[high] {
			arr[higherThanIndex], arr[unknownIndex] = arr[unknownIndex], arr[higherThanIndex]
			higherThanIndex++
		}

		unknownIndex++
	}

	// recursively sort the smaller and higher than pivot sub arrays
	left, right := QuickSort(arr[:higherThanIndex]), QuickSort(arr[higherThanIndex:])
	return append(left, right...)
}
