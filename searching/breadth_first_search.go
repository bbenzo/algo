package searching

type Node struct {
	Index       int
	Predecessor *Node
	Distance    int
	Neighbors   []*Node
}

func BreadthFirstSearch(root *Node) {
	for _, neighbor := range root.Neighbors {
		if neighbor.Distance != 0 {
			continue
		}

		neighbor.Predecessor = root
		neighbor.Distance = root.Distance + 1

		BreadthFirstSearch(neighbor)
	}
}
