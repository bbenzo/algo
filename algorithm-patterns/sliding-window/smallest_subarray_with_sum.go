package sliding_window

const Infinity = int(^uint(0) >> 1)

// Given an array of positive numbers and a positive number ‘S’, find the length of the smallest contiguous subarray
// whose sum is greater than or equal to ‘S’. Return 0, if no such subarray exists.
func LengthOfSmallestContiguousSubArrayWithSum(arr []int, s int) int {
	result, windowStart, sum := Infinity, 0, 0
	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		sum += arr[windowEnd]

		for sum >= s {
			partialResult := windowEnd - windowStart + 1
			if partialResult < result {
				result = partialResult
			}

			sum -= arr[windowStart]
			windowStart++
		}
	}

	if result == Infinity {
		return 0
	}

	return result
}
