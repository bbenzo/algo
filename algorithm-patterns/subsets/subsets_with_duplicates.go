package subsets

import "sort"

func SubsetsWithDuplicates(arr []int) [][]int {
	result := [][]int{}
	result = append(result, []int{})

	sort.Ints(arr)

	for i := range arr {

		length := len(result)
		j := 0

		if i > 0 && arr[i] == arr[i-1] {
			j = length / 2
		}

		for j < length {
			subset := result[j]
			newSubset := append([]int{}, subset...)
			newSubset = append(newSubset, arr[i])
			result = append(result, newSubset)

			j++
		}

	}

	return result
}
