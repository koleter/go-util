package list

import (
	"github.com/koleter/go-util/concurrency/lock"
	"sort"
)

type threadSafeList[T any] struct {
	lock *lock.ReentrantMutex
	list []T
}

func NewThreadSafeList[T any](l []T) *threadSafeList[T] {
	return &threadSafeList[T]{
		lock: new(lock.ReentrantMutex),
		list: l,
	}
}

func (tsl *threadSafeList[T]) WithLock(f func()) {
	tsl.lock.Lock()
	defer tsl.lock.Unlock()
	f()
}

func (tsl *threadSafeList[T]) Append(element ...T) {
	tsl.lock.Lock()
	defer tsl.lock.Unlock()
	tsl.list = append(tsl.list, element...)
}

func (tsl *threadSafeList[T]) Get(i int) T {
	tsl.lock.Lock()
	defer tsl.lock.Unlock()
	return tsl.list[i]
}

func (tsl *threadSafeList[T]) Set(i int, element T) T {
	tsl.lock.Lock()
	defer tsl.lock.Unlock()
	old := tsl.list[i]
	tsl.list[i] = element
	return old
}

func (tsl *threadSafeList[T]) Remove(index int) T {
	tsl.lock.Lock()
	defer tsl.lock.Unlock()
	removed := tsl.list[index]
	tsl.list = append(tsl.list[:index], tsl.list[index+1:]...)
	return removed
}

func (tsl *threadSafeList[T]) RemoveFunc(f func(t T) bool) []T {
	var res []T
	tsl.lock.Lock()
	defer tsl.lock.Unlock()
	for i := len(tsl.list) - 1; i >= 0; i-- {
		if f(tsl.list[i]) {
			res = append(res, tsl.list[i])
			tsl.list = append(tsl.list[:i], tsl.list[i+1:]...)
		}
	}
	return res
}

// Len 返回列表长度
func (tsl *threadSafeList[T]) Len() int {
	return len(tsl.list)
}

func (tsl *threadSafeList[T]) Range(f func(int, T) bool) {
	tsl.lock.Lock()
	defer tsl.lock.Unlock()
	for i, t := range tsl.list {
		if !f(i, t) {
			return
		}
	}
}

// Contain 是否存在满足要求的元素
func (tsl *threadSafeList[T]) Contain(f func(int, T) bool) bool {
	tsl.lock.Lock()
	defer tsl.lock.Unlock()
	for i, t := range tsl.list {
		if f(i, t) {
			return true
		}
	}
	return false
}

func (tsl *threadSafeList[T]) Filter(f func(int, T) bool) []T {
	var ret []T
	tsl.lock.Lock()
	defer tsl.lock.Unlock()
	for i, t := range tsl.list {
		if f(i, t) {
			ret = append(ret, t)
		}
	}
	return ret
}

func (tsl *threadSafeList[T]) Clear() {
	tsl.lock.Lock()
	defer tsl.lock.Unlock()
	tsl.list = tsl.list[:0]
}

func (tsl *threadSafeList[T]) Sort(f func(i, j int) bool) {
	tsl.lock.Lock()
	defer tsl.lock.Unlock()
	sort.Slice(tsl.list, f)
}
