package fast_slow_pointers

func CycleCircularArray(arr []int) bool {
	if len(arr) == 0 {
		return false
	}

	slow, fast := 0, 0
	for slow < len(arr) {
		slow += arr[slow]
		if slow >= len(arr) {
			slow -= len(arr)
		}

		fast += arr[fast]
		if fast >= len(arr) {
			fast -= len(arr)
		}

		fast += arr[fast]
		if fast >= len(arr) {
			fast -= len(arr)
		}

		if slow == fast {
			return true
		}
	}

	return false
}
