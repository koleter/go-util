package dlinkedlist

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestDoublyLinkedList_PopBack(t *testing.T) {
	d := DoublyLinkedList[int]{}
	d.PushFront(3)
	assert.Equal(t, 1, d.Len())
	d.PushBack(2)
	assert.Equal(t, 2, d.Len())
	back, exist := d.PopBack()
	assert.True(t, exist)
	assert.Equal(t, 2, back)
}

func TestDoublyLinkedList_PopFront(t *testing.T) {
	d := DoublyLinkedList[int]{}
	back, exist := d.PopFront()
	assert.False(t, exist)
	d.PushFront(3)
	d.PushBack(2)
	back, exist = d.PopFront()
	assert.True(t, exist)
	assert.Equal(t, 3, back)
}

func TestDoublyLinkedList_for_range(t *testing.T) {
	d := DoublyLinkedList[int]{}
	for i := 0; i < 10; i++ {
		d.PushBack(i)
	}
	for i := 0; i < 10; i++ {
		back, exist := d.PopFront()
		assert.True(t, exist)
		assert.Equal(t, i, back)
	}
}

func TestNode_Remove(t *testing.T) {
	d := DoublyLinkedList[int]{}
	for i := 0; i < 10; i++ {
		d.PushBack(i)
	}
	randInt := rand.Intn(10)
	assert.Equal(t, 10, d.Len())

	for node := d.Head(); node != nil; node = node.Next() {
		if node.Value == randInt {
			d.Remove(node)
		}
	}
	assert.Equal(t, 9, d.Len())

	for node := d.Head(); node != nil; node = node.Next() {
		assert.False(t, node.Value == randInt)
	}
}

func TestNode_Remove_All(t *testing.T) {
	d := DoublyLinkedList[int]{}
	for i := 0; i < 10; i++ {
		d.PushBack(i)
	}

	removed := 0
	for node := d.Head(); node != nil; node = node.Next() {
		assert.Equal(t, 10-removed, d.Len())
		d.Remove(node)
		removed++
	}
	assert.Equal(t, 0, d.Len())
}

func TestNode_Remove_one_element_list(t *testing.T) {
	d := DoublyLinkedList[int]{}

	front := d.PushFront(3)
	d.Remove(front)
}
