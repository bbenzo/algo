package two_heaps

// bruteforce solution: keep a sorted list and insert in the middle -> O(n)
// two heap solution: x = median. divide list into two halves. one for all smaller than x = smallerList.
// one for all greater than x = greaterList. median is either largest of smallerList or smallest of greaterList.
// if number of elements is even, median is average of both.

type stream struct {
	minHeap *MinIntHeap
	maxHeap *MaxIntHeap
}

func newStream() *stream {
	return &stream{
		minHeap: &MinIntHeap{},
		maxHeap: &MaxIntHeap{},
	}
}

func (s *stream) insert(val int) {
	if len(*s.maxHeap) == 0 || val < (*s.maxHeap)[0] {
		s.maxHeap.Push(val)
	} else {
		s.minHeap.Push(val)
	}

	diff := s.maxHeap.Len() - s.minHeap.Len()
	if diff > 1 {
		largest := s.maxHeap.Pop()
		s.minHeap.Push(largest)
	} else if diff < -1 {
		smallest := s.minHeap.Pop()
		s.maxHeap.Push(smallest)
	}
}

func (s *stream) median() float64 {
	if s.maxHeap.Len() == 0 {
		return -1
	}

	diff := s.maxHeap.Len() - s.minHeap.Len()
	if diff == 0 {
		return (float64(s.minHeap.Pop()) + float64(s.maxHeap.Pop())) / 2
	} else if diff > 0 {
		return float64(s.maxHeap.Pop())
	} else {
		return float64(s.minHeap.Pop())
	}
}
