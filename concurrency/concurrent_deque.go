package concurrency

import (
	"github.com/koleter/go-util/concurrency/lock"
	"github.com/koleter/go-util/queue"
)

type ConcurrentDeque[T any] struct {
	lock  *lock.ReentrantMutex
	deque queue.Deque[T]
}

func NewConcurrentDeque[T any](deque queue.Deque[T]) *ConcurrentDeque[T] {
	return &ConcurrentDeque[T]{
		lock:  new(lock.ReentrantMutex),
		deque: deque,
	}
}

func (c *ConcurrentDeque[T]) WithLock(f func()) {
	c.lock.Lock()
	defer c.lock.Unlock()
	f()
}

func (c *ConcurrentDeque[T]) IsEmpty() bool {
	return c.deque.IsEmpty()
}

func (c *ConcurrentDeque[T]) IsFull() bool {
	return c.deque.IsFull()
}

func (c *ConcurrentDeque[T]) Size() int {
	return c.deque.Size()
}

func (c *ConcurrentDeque[T]) Capacity() int {
	return c.deque.Capacity()
}

func (c *ConcurrentDeque[T]) PushFront(item T) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.deque.PushFront(item)
}

func (c *ConcurrentDeque[T]) PushBack(item T) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.deque.PushBack(item)
}

func (c *ConcurrentDeque[T]) PopFront() (T, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.deque.PopFront()
}

func (c *ConcurrentDeque[T]) PopBack() (T, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.deque.PopBack()
}

func (c *ConcurrentDeque[T]) Front() (T, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.deque.Front()
}

func (c *ConcurrentDeque[T]) Back() (T, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.deque.Back()
}

func (c *ConcurrentDeque[T]) Range(fn func(T) bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.deque.Range(fn)
}

func (c *ConcurrentDeque[T]) ReverseRange(fn func(item T) bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.ReverseRange(fn)
}

func (c *ConcurrentDeque[T]) Clear() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.deque.Clear()
}
