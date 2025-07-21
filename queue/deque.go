package queue

type Deque[T any] interface {
	IsEmpty() bool
	IsFull() bool
	Size() int
	Capacity() int
	PushFront(item T)
	PushBack(item T)
	PopFront() (T, bool)
	PopBack() (T, bool)
	Front() (T, bool)
	Back() (T, bool)
	Range(fn func(T) bool)
	ReverseRange(fn func(item T) bool)
	Clear()
}
