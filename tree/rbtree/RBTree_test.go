package rbtree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Key int

func (k Key) Compare(other Key) int {
	if k > other {
		return 1
	} else if k == other {
		return 0
	}
	return -1
}

func TestRBTree_Delete(t *testing.T) {
	tree := RBTree[Key, int]{}
	tree.Insert(3, 5)
	get, exist := tree.Get(3)
	assert.True(t, exist)
	assert.Equal(t, 5, get)

	tree.Delete(3)
	_, exist = tree.Get(3)
	assert.False(t, exist)
}

func TestRBTree_Higher(t *testing.T) {
	tree := RBTree[Key, int]{}
	for i := 0; i < 20; i++ {
		tree.Insert(Key(i), i)
		assert.Equal(t, i+1, tree.Len())
	}

	higher, exist := tree.Higher(5)
	assert.True(t, exist)
	assert.Equal(t, 6, higher)
}

func TestRBTree_Lower(t *testing.T) {
	tree := RBTree[Key, int]{}
	for i := 0; i < 20; i++ {
		tree.Insert(Key(i), i)
		assert.Equal(t, i+1, tree.Len())
	}

	lower, exist := tree.Lower(13)
	assert.True(t, exist)
	assert.Equal(t, 12, lower)
}

func TestRBTree_Len(t *testing.T) {
	tree := RBTree[Key, int]{}
	for i := 0; i < 20; i++ {
		tree.Insert(Key(i), i)
		assert.Equal(t, i+1, tree.Len())
	}

	// insert a exist key, length should not change
	tree.Insert(2, 2)
	assert.Equal(t, 20, tree.Len())
}
