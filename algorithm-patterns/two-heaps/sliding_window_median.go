package two_heaps

func SlidingWindowMedian(nums []int, k int) []float64 {
	var result []float64
	s := newStream()

	i := 0
	for i < len(nums) {
		s.insert(nums[i])

		j := i - k
		if j > -1 {
			s.remove(nums[j])
		}

		if i >= k-1 {
			result = append(result, s.median())
		}

		i++
	}

	return result
}
