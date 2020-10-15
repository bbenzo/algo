package graphs

type AdjacencyList [][]int

func SampleAdjacencyList() AdjacencyList {
	return AdjacencyList{
		[]int{1, 2, 5},
		[]int{1},
		[]int{4},
		[]int{3, 5},
		[]int{1, 3, 4},
	}
}