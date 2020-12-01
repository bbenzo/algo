package two_pointers

func SumOfPairEqualSum(arr []int, targetSum int) (int, int) {
	start := 0
	end := len(arr) - 1

	for end > start {
		sum := arr[start] + arr[end]
		if sum == targetSum {
			return start, end
		} else if sum < targetSum {
			start++
		} else {
			end--
		}
	}

	return -1, -1
}
