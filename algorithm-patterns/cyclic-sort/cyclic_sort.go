package cyclic_sort

func CyclicSort(arr []int) {
	i := 0
	for i < len(arr) {
		if arr[i] != i+1 {
			arr[i], arr[arr[i]-1] = arr[arr[i]-1], arr[i]
		} else {
			i++
		}
	}
}
