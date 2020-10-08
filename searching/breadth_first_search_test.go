package searching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
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

	BreadthFirstSearch(root)

	assert.Equal(t, root.Distance, 0)
	assert.Equal(t, root.Neighbors[0].Distance, 1)
	assert.Equal(t, root.Neighbors[0].Neighbors[0].Distance, 2)
	assert.Equal(t, root.Neighbors[0].Neighbors[0].Neighbors[0].Distance, 3)
	assert.Equal(t, root.Neighbors[0].Neighbors[0].Neighbors[0].Neighbors[0].Distance, 4)
	assert.Nil(t, root.Neighbors[0].Neighbors[0].Neighbors[0].Neighbors[0].Neighbors)

	assert.Equal(t, root.Neighbors[0].Neighbors[1].Distance, 2)
	assert.Equal(t, root.Neighbors[0].Neighbors[1].Neighbors[0].Distance, 3)
	assert.Nil(t, root.Neighbors[0].Neighbors[1].Neighbors[0].Neighbors)

	assert.Equal(t, root.Neighbors[1].Distance, 1)
	assert.Equal(t, root.Neighbors[1].Neighbors[0].Distance, 2)
	assert.Equal(t, root.Neighbors[1].Neighbors[0].Neighbors[0].Distance, 3)
	assert.Nil(t, root.Neighbors[1].Neighbors[0].Neighbors[0].Neighbors)

}
