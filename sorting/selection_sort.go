package sorting

func SelectionSort(list []int32) []int32 {
	for i := 0; i < len(list); i++ {
		for j := i; j < len(list); j++ {
			if list[j] < list[i] {
				swap := list[i]
				list[i] = list[j]
				list[j] = swap
			}
		}
	}
	return list
}
