package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (heap IntHeap) Len() int {
	return len(heap)
}

func (heap IntHeap) Less(i, j int) bool {
	return heap[i] < heap[j]
}

func (heap IntHeap) Swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func (heap *IntHeap) Push(heapInterface interface{}) {
	*heap = append(*heap, heapInterface.(int))
}

func (heap *IntHeap) Pop() interface{} {
	var previous IntHeap = *heap
	n := len(previous)
	x1 := previous[n-1]
	*heap = previous[0 : n-1]
	return x1
}

func main() {
	var intHeap *IntHeap = &IntHeap{1, 4, 7}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	heap.Push(intHeap, 6)

	fmt.Printf("Minimum: %d \n", (*intHeap)[0])

	for intHeap.Len() > 0 {
 	 	fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}
