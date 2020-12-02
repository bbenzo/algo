package two_pointers

import "sort"

func TripletsSmallerSum(arr []int, target int) int {
	sort.Ints(arr)

	count := 0

	for i := 0; i < len(arr); i++ {
		left, right := i+1, len(arr) - 1

		for left < right {
			sum := arr[left] + arr[right] + arr[i]

			if sum < target {
				count += right - left // every number for right will now be viable
				left++
			} else {
				right--
			}
		}
	}

	return count
}
