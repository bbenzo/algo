package dijkstras

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortestPath(t *testing.T) {
	graph := &Graph{}

	a := &Node{Name: "A"}
	graph.AddNode(a)
	b := &Node{Name: "B"}
	graph.AddNode(b)
	c := &Node{Name: "C"}
	graph.AddNode(c)
	d := &Node{Name: "D"}
	graph.AddNode(d)

	graph.AddEdge(a, b, 7)
	graph.AddEdge(a, b, 3)
	graph.AddEdge(b, c, 2)
	graph.AddEdge(c, d, 1)

	result := Dijkstra(a, d, graph)

	assert.NotNil(t, result)
	assert.Equal(t, 3, result[b])
	assert.Equal(t, 5, result[c])
	assert.Equal(t, 6, result[d])
}
