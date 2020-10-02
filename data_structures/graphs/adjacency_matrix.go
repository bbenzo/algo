package graphs

type AdjacencyMatrix [][]int

func SampleAdjacencyMatrix() AdjacencyMatrix {
	return AdjacencyMatrix{
		[]int{0, 1, 1},
		[]int{1, 0, 1},
		[]int{1, 1, 0},
	}
}