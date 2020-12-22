package merge_intervals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConflictingAppointments(t *testing.T) {
	input := Intervals{{1,4}, {2,5}, {7,9}}

	assert.False(t, ConflictingAppointments(input))
}

func TestConflictingAppointments2(t *testing.T) {
	input := Intervals{{6,7}, {2,4}, {8,12}}

	assert.True(t, ConflictingAppointments(input))
}

func TestConflictingAppointments3(t *testing.T) {
	input := Intervals{{4,5}, {2,3}, {3,6}}

	assert.False(t, ConflictingAppointments(input))
}
