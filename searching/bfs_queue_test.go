package searching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBFSQueue(t *testing.T) {
		root := &Node{
		Index:       0,
		Predecessor: nil,
		Distance:    0,
		Neighbors:   []*Node{
			{
				Index: 1,
				Neighbors: []*Node{
					{
						Index: 3,
						Neighbors: []*Node{
							{
								Index: 5,
								Neighbors: []*Node{
									{
										Index: 8,
									},
								},
							},
						},
					},
					{
						Index: 4,
						Neighbors: []*Node{
							{
								Index: 7,
							},
						},
					},
				},
			},
			{
				Index: 2,
				Neighbors: []*Node{
					{
						Index: 5,
						Neighbors: []*Node{
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
