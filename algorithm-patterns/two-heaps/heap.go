package two_heaps

type MinIntHeap []int

func (h MinIntHeap) Len() int           { return len(h) }
func (h MinIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinIntHeap) Push(x int) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x)
}

func (h *MinIntHeap) Pop() int {
	old := *h
	x := old[0]
	*h = old[1:]
	return x
}

func (h *MinIntHeap) Remove(val int) {
	index := -1

	for i := range *h {
		if (*h)[i] == val {
			index = i
			break
		}
	}

	if index == -1 {
		return
	}

	*h = append((*h)[:index], (*h)[index+1:]...)
}

type MaxIntHeap []int

func (h MaxIntHeap) Len() int           { return len(h) }
func (h MaxIntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxIntHeap) Push(x int) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x)
}

func (h *MaxIntHeap) Pop() int {
	old := *h
	x := old[0]
	*h = old[1:]
	return x
}

func (h *MaxIntHeap) Top() int {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *MaxIntHeap) Remove(val int) {
	index := -1

	for i := range *h {
		if (*h)[i] == val {
			index = i
			break
		}
	}

	if index == -1 {
		return
	}

	*h = append((*h)[:index], (*h)[index+1:]...)
}
