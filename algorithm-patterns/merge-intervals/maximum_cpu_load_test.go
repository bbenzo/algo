package merge_intervals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximumCPULoad1(t *testing.T) {
	jobs := Jobs{{1, 4, 3}, {2, 5, 4}, {7, 9, 6}}

	assert.Equal(t, 7, MaximumCPULoad(jobs))
}

func TestMaximumCPULoad2(t *testing.T) {
	jobs := Jobs{{6, 7, 10}, {2, 4, 11}, {8, 12, 15}}

	assert.Equal(t, 15, MaximumCPULoad(jobs))
}


func TestMaximumCPULoad3(t *testing.T) {
	jobs := Jobs{{1, 4, 2}, {2, 4, 1}, {3, 6, 5}}

	assert.Equal(t, 8, MaximumCPULoad(jobs))
}