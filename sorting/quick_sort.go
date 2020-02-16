package sorting

func QuickSort(array []int) []int {
	n := len(array)
	pivotIndex := n - 1
	return partition(array, pivotIndex, 0, pivotIndex - 1)
}

func partition(arr []int, pivotIdx int, leftIdx int, rightIdx int) []int {
	if len(arr) < 2 {
		return arr
	}

	foundLeft := false
	for leftIdx < rightIdx && leftIdx < pivotIdx {
		if arr[leftIdx] > arr[pivotIdx] {
			foundLeft = true
			break
		}
		leftIdx++
	}

	foundRight := false
	for rightIdx >= leftIdx && rightIdx > -1 {
		if arr[rightIdx] < arr[pivotIdx] {
			foundRight = true
			break
		}
		rightIdx--
	}

	if foundLeft && foundRight {
		swap(arr, leftIdx, rightIdx)
		partition(arr, pivotIdx, leftIdx, rightIdx)
	}

	if leftIdx > rightIdx {
		swap(arr, leftIdx, pivotIdx)

		partition1 := arr[0:leftIdx]
		pivotIndex1 := len(partition1) - 1
		partition2 := arr[leftIdx:]
		pivotIndex2 := len(partition2) - 1

		partition(partition1, pivotIndex1, 0, pivotIndex1-1)
		partition(partition2, pivotIndex2, 0, pivotIndex2-1)
	}
	return arr
}

func swap(arr []int, leftIdx int, rightIdx int) {
	swap := arr[leftIdx]
	arr[leftIdx] = arr[rightIdx]
	arr[rightIdx] = swap
}
