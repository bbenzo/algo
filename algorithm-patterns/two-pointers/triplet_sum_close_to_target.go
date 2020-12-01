package two_pointers

import (
	"sort"
)

func TripletSumCloseToTarget(arr []int, target int) int {
	sort.Ints(arr)

	result := 9999999999
	resultDiff := 99999999

	for i := 0; i < len(arr); i++ {
		left, right := i+1, len(arr) - 1

		for left < right {
			sum := arr[left] + arr[right] + arr[i]
			diff := sum - target

			if Abs(diff) < resultDiff {
				resultDiff = Abs(diff)
				result = sum
			}

			if diff < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
