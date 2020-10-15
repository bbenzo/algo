package breadth_first_search

import (
	"github.com/stretchr/testify/assert"
	queue2 "github/bbenzo/algo/data-structures/queue"
	"testing"
)

func TestBFSQueue(t *testing.T) {
		root := &queue2.Node{
		Index:       0,
		Predecessor: nil,
		Distance:    0,
		Neighbors:   []*queue2.Node{
			{
				Index: 1,
				Neighbors: []*queue2.Node{
					{
						Index: 3,
						Neighbors: []*queue2.Node{
							{
								Index: 5,
								Neighbors: []*queue2.Node{
									{
										Index: 8,
									},
								},
							},
						},
					},
					{
						Index: 4,
						Neighbors: []*queue2.Node{
							{
								Index: 7,
							},
						},
					},
				},
			},
			{
				Index: 2,
				Neighbors: []*queue2.Node{
					{
						Index: 5,
						Neighbors: []*queue2.Node{
							{
								Index: 6,
							},
						},
					},
				},
			},
		},
	}

	result := BreadthFirstSearchQueue(root, 8)

	assert.NotNil(t, result)
	assert.Equal(t, 8, result.Index)
	assert.Equal(t, 4, result.Distance)
	assert.Equal(t, 5, result.Predecessor.Index)
	assert.Equal(t, 3, result.Predecessor.Predecessor.Index)
	assert.Equal(t, 1, result.Predecessor.Predecessor.Predecessor.Index)
}

func BenchmarkBreadthFirstSearchQueue(b *testing.B) {
	root := &queue2.Node{
		Index:       0,
		Predecessor: nil,
		Distance:    0,
		Neighbors:   []*queue2.Node{
			{
				Index: 1,
				Neighbors: []*queue2.Node{
					{
						Index: 3,
						Neighbors: []*queue2.Node{
							{
								Index: 5,
								Neighbors: []*queue2.Node{
									{
										Index: 8,
									},
								},
							},
						},
					},
					{
						Index: 4,
						Neighbors: []*queue2.Node{
							{
								Index: 7,
							},
						},
					},
				},
			},
			{
				Index: 2,
				Neighbors: []*queue2.Node{
					{
						Index: 5,
						Neighbors: []*queue2.Node{
							{
								Index: 6,
							},
						},
					},
				},
			},
		},
	}

	for n := 0; n < b.N; n++ {
		BreadthFirstSearchQueue(root, 8)
	}
}
