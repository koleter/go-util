package list

import "github.com/koleter/go-util/concurrency/lock"

type List[T any] interface {
	Append(element ...T)
	Get(i int) T
	Set(i int, element T) T
	Remove(index int) T
	RemoveFunc(f func(t T) bool) []T
	Len() int
	Range(f func(int, T) bool)
	Contain(f func(int, T) bool) bool
	Filter(f func(int, T) bool) []T
	Clear()
	Sort(f func(i, j int) bool)
}

type ConcurrentList[T any] interface {
	lock.Locker
	List[T]
}
