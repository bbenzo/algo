package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitNewWithElements(t *testing.T) {
	list := New(2)

	assert.Equal(t,  1, len(list))
	assert.Equal(t, 2, list[0].Data, )
}

func TestAppendToEmptySlice(t *testing.T) {
	empty := New()
	filled := Append(empty,2)

	assert.Equal(t, 1, len(filled))
	assert.Equal(t, 2, filled[0].Data)
}

func TestAppendOnPrefilled(t *testing.T) {
	empty := New(4)
	filled := Append(empty,2)

	assert.Equal(t, 2, len(filled))
	assert.Equal(t, 4, filled[0].Data)
	assert.Equal(t, 2, filled[1].Data)
	assert.Equal(t, 2, filled[0].Next.Data)
	assert.Equal(t, 4, filled[1].Last.Data)
}

func TestRemove(t *testing.T) {
	list := New(4, 5, 6)
	list = Remove(list, 1)

	assert.Equal(t, 2, len(list))
	assert.Equal(t, 6, list[0].Next.Data)
	assert.Equal(t, 4, list[1].Last.Data)
}

func TestRemoveFromOneItemList(t *testing.T) {
	list := New(4)
	list = Remove(list, 0)

	assert.Equal(t, 0, len(list))
}

func TestRemoveOutOfBoundIndex(t *testing.T) {
	list := New(2, 3)
	list = Remove(list, 2)

	assert.Equal(t, 2, len(list))
}

func BenchmarkNewListWith10Items(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		New(26, 8, 2, 5, 6, 7, 7, 3 ,3 ,2)
	}
}

func BenchmarkInsertItem(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		list := New(26, 8, 2, 5, 6, 7, 7, 3 ,3 ,2)
		list = Insert(list,5, 5)
	}
}

func BenchmarkRemoveItem(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		list := New(26, 8, 2, 5, 6, 7, 7, 3 ,3 ,2)
		Remove(list, 5)
	}
}
