package graphs

/*
First element: vertex one
Second element: vertex two
Third element: edge weight
 */

type EdgeList [][3]int

func SampleEdgeList() EdgeList {
	return EdgeList{
		[3]int{1, 2, 4},
		[3]int{2, 3, 7},
		[3]int{3, 1, 2},
	}
}
