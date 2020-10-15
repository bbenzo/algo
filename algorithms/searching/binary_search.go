package searching

// BinarySearch returns position of item in list
func BinarySearch(a int, sortedList []int) int {
	left := 0
	right := len(sortedList) - 1

	for left <= right {
		i := right + left / 2

		if a == sortedList[i] {
			return i
		} else if a > sortedList[i] {
			left = i + 1
		} else {
			right = i - 1
		}
	}
	return -1
}
