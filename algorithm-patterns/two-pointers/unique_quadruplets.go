package two_pointers

import "sort"

func UniqueQuarduplets(arr []int, target int) [][]int {
	result := [][]int{}

	sort.Ints(arr)

	for i := range arr {
		start, end := i+1, len(arr) - 1

		for start < end {
			sum := arr[i] + arr[start] + arr[end]

			k := start + 1
			for k < end {
				if sum + arr[k] == target {
					result = append(result, []int{arr[i], arr[start], arr[k], arr[end]})
				}

				for arr[k] == arr[k+1] {
					k++
				}

				k++
			}

			if sum < target {
				start++
			} else {
				end--
			}
		}
	}

	return result
}
