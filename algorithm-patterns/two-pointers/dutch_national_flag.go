package two_pointers

func DutchNationalFlag(arr []int) []int {
	low, high := 0, len(arr) - 1

	i := 0
	for i <= high {
		switch arr[i] {
		case 0:
			arr[i], arr[low] = arr[low], arr[i]
			low++
			i++
		case 1:
			i++
		case 2:
			arr[i], arr[high] = arr[high], arr[i]
			high--
		}
	}

	return arr
}
