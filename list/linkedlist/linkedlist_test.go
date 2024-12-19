package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_Len(t *testing.T) {
	var linklist = &LinkedList[int]{}
	linklist.Append(3)
	linklist.Append(31)
	assert.Equal(t, 2, linklist.Len())
	pop, _ := linklist.Pop()
	assert.Equal(t, 3, pop)
	pop, _ = linklist.Pop()
	assert.Equal(t, 31, pop)
	pop, exist := linklist.Pop()
	assert.False(t, exist)
	assert.Equal(t, 0, linklist.Len())
}
