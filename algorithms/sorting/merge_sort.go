package sorting

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := len(arr) / 2
	return merge(MergeSort(arr[pivot:]), MergeSort(arr[:pivot]))
}

func merge(arr1, arr2 []int) []int {
	result := make([]int, len(arr1)+len(arr2))

	count1, count2 := 0, 0
	for i := range result {
		if count1 >= len(arr1) {
			result[i] = arr2[count2]
			count2++
		} else if count2 >= len(arr2) || arr1[count1] <= arr2[count2] {
			result[i] = arr1[count1]
			count1++
		} else {
			result[i] = arr2[count2]
			count2++
		}
	}
	return result
}
