package two_heaps

func SlidingWindowMedian(nums []int, k int) []float64 {
	var result []float64

	i := k
	for i <= len(nums) {
		j := i - k

		s := newStream()
		for j < i {
			s.insert(nums[j])
			j++
		}

		result = append(result, s.median())

		i++
	}

	return result
}
