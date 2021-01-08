package two_heaps

func MaximizeCapital(capitals, profits []int, initialCapital, projectNums int) int {
	minHeap := &MinIntHeap{}
	maxHeap := &MaxIntHeap{}

	lastAdded := 0
	for projectNums > 0 {
		for i := lastAdded; i < len(capitals); i++ {
			if capitals[i] <= initialCapital {
				minHeap.Push(capitals[i])
				maxHeap.Push(profits[i])

				lastAdded = i
			}
		}

		max := maxHeap.Top()
		initialCapital += max

		projectNums--
	}

	return initialCapital
}
