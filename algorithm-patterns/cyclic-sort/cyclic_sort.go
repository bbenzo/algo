package cyclic_sort

func CyclicSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] != i+1 {
			arr[i], arr[arr[i]-1] = arr[arr[i]-1], arr[i]
		}
	}
}
