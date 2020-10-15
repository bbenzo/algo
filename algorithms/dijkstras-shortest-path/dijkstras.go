package dijkstras

import (
	"strconv"
)

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

type Edge struct {
	Start  *Node
	End    *Node
	Weight int
}

type Node struct {
	Name string
}

const Infinity = int(^uint(0) >> 1)

func (g *Graph) AddEdge(start *Node, end *Node, weight int) {
	g.Edges = append(g.Edges, &Edge{
		Start:  start,
		End:    end,
		Weight: weight,
	})
}

func (g *Graph) AddNode(node *Node) {
	for i := range g.Nodes {
		if g.Nodes[i] == node {
			return
		}
	}

	g.Nodes = append(g.Nodes, node)
}

func (g *Graph) String() string {
	var s string

	s += "Edges:\n"
	for _, edge := range g.Edges {
		s += edge.Start.Name + " -> " + edge.End.Name + " = " + strconv.Itoa(edge.Weight)
		s += "\n"
	}
	s += "\n"

	s += "Nodes: "
	for i, node := range g.Nodes {
		if i == len(g.Nodes)-1 {
			s += node.Name
		} else {
			s += node.Name + ", "
		}
	}
	s += "\n"

	return s
}

func Dijkstra(start *Node, end *Node, graph *Graph) map[*Node]int {
	distancesToStart := initDistancesToStart(start, graph)

	q := initQueue(graph)

	visitedNodes := make(map[*Node]bool)
	visitedNodes[start] = true

	for q.Len() > 0 {
		next := nextNodeWithMinDistance(q, distancesToStart, visitedNodes)

		setVisited(visitedNodes, next)

		for _, node := range graph.Nodes {
			if !visitedNodes[node] {
				q.Enqueue(node)
			}
		}

		nodeEdges := edges(next, graph)

		for _, edge := range nodeEdges {
			combinedWeight := distancesToStart[next] + edge.Weight

			if combinedWeight < distancesToStart[edge.End] {
				distancesToStart[edge.End] = combinedWeight
			}
		}
	}

	return distancesToStart
}

func initQueue(graph *Graph) Queue {
	q := NewQueue()
	for _, node := range graph.Nodes {
		q.Enqueue(node)
	}
	return q
}

func initDistancesToStart(start *Node, graph *Graph) map[*Node]int {
	distancesToStart := make(map[*Node]int)
	distancesToStart[start] = 0
	for _, node := range graph.Nodes {
		if node != start {
			distancesToStart[node] = Infinity
		}
	}
	return distancesToStart
}

func setVisited(visitedNodes map[*Node]bool, next *Node) {
	visitedNodes[next] = true
}

func nextNodeWithMinDistance(q Queue, distancesToStart map[*Node]int, visitedNodes map[*Node]bool) *Node {
	next, _ := q.Dequeue()

	for q.Len() > 0 {
		item, _ := q.Dequeue()
		if !visitedNodes[next] && distancesToStart[item] < distancesToStart[next] {
			next = item
		}
	}
	return next
}

func edges(node *Node, g *Graph) []*Edge {
	var nodeEdges []*Edge

	for _, edge := range g.Edges {
		if edge.Start == node {
			nodeEdges = append(nodeEdges, edge)
		}
	}

	return nodeEdges
}
