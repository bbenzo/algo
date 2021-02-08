package modified_binary_search

import "math"

// assume we have an array that has no size

type arrReader []int

func (a arrReader) Get(i int) int {
	if i >= len(a) {
		return math.MaxInt64
	}

	return a[i]
}

func SearchSortedInfiniteArray(arr arrReader, key int) int {
	left, right := 0, 1
	for right < math.MaxInt64 {
		if key <= arr[right] {
			break
		} else {
			left = (left + 1) * 2
			right = (right + 1) * 2 + left
		}
	}


	for left <= right {
		pivot := (left + right) / 2

		if arr[pivot] == key {
			return pivot
		}

		if key < arr[pivot] {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

	return -1
}
