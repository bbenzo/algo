package two_pointers

func TriplestWithProductLessTarget(arr []int, target int) [][]int {
	result := [][]int{}

	for i := 0; i < len(arr); i++ {
		product := 1

		j := i
		tempArr := []int{}
		for product < target && j < len(arr) {
			product *= arr[j]

			if product < target {
				tempArr = append(tempArr, arr[j])
				result = append(result, tempArr)
			}

			j++
		}
	}

	return result
}
