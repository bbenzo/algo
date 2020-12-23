package merge_intervals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmployeeFreeTime1(t *testing.T) {
	employeeMeetings := []Meetings{{{1, 3}, {5, 6}}, {{2, 3}, {6, 8}}}

	expected := Meetings{{3, 5}}

	assert.Equal(t, expected, EmployeeFreeTime(employeeMeetings))
}

func TestEmployeeFreeTime2(t *testing.T) {
	employeeMeetings := []Meetings{{{1, 3}, {9, 12}}, {{2, 4}, {6, 8}}}

	expected := Meetings{{4, 6}, {8, 9}}

	assert.Equal(t, expected, EmployeeFreeTime(employeeMeetings))
}

func TestEmployeeFreeTime3(t *testing.T) {
	employeeMeetings := []Meetings{{{1, 3}, {2, 4}}, {{3, 5}, {7, 9}}}

	expected := Meetings{{5, 7}}

	assert.Equal(t, expected, EmployeeFreeTime(employeeMeetings))
}
