package fast_slow_pointers

func CycleCircularArray(arr []int) bool {
	if len(arr) == 0 {
		return false
	}

	slow, fast := 0, 0
	for slow < len(arr) {
		oneDirection := true

		slow += arr[slow]
		if slow >= len(arr) {
			slow -= len(arr)
		}

		slowToTheRight := goesToTheRight(arr, slow)

		fast += arr[fast]
		if fast >= len(arr) {
			fast -= len(arr)
		}

		fastOneToTheRight := goesToTheRight(arr, fast)

		fast += arr[fast]
		if fast >= len(arr) {
			fast -= len(arr)
		}

		fastTwoToTheRight := goesToTheRight(arr, fast)

		if !slowToTheRight || !fastOneToTheRight || !fastTwoToTheRight {
			oneDirection = false
		}

		if !slowToTheRight && !fastOneToTheRight && !fastTwoToTheRight {
			oneDirection = true
		}

		if slow == fast && oneDirection {
			return true
		} else if slow == fast && !oneDirection {
			return false
		}
	}

	return false
}

func goesToTheRight(arr []int, slow int) bool {
	slowToTheRight := true
	if arr[slow] < 0 {
		slowToTheRight = false
	}
	return slowToTheRight
}
