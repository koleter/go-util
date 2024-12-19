package _map

import (
	"github.com/koleter/go-util/concurrency/lock"
)

type ThreadSafeMap[K comparable, V any] struct {
	lock    *lock.ReentrantMutex
	raw_map map[K]V
}

func NewThreadSafeMap[K comparable, V any](raw_map map[K]V) *ThreadSafeMap[K, V] {
	if raw_map == nil {
		panic("can not use nil map to new ThreadSafeMap")
	}
	return &ThreadSafeMap[K, V]{
		lock:    new(lock.ReentrantMutex),
		raw_map: raw_map,
	}
}

func (t *ThreadSafeMap[K, V]) WithLock(f func()) {
	t.lock.Lock()
	defer t.lock.Unlock()
	f()
}

func (t *ThreadSafeMap[K, V]) Put(key K, val V) V {
	t.lock.Lock()
	defer t.lock.Unlock()
	v := t.raw_map[key]
	t.raw_map[key] = val
	return v
}

func (t *ThreadSafeMap[K, V]) PutAll(m map[K]V) {
	t.lock.Lock()
	defer t.lock.Unlock()
	for k, v := range m {
		t.raw_map[k] = v
	}
}

// PutIfAbsent 只有不存在相同的key时才会保存
func (t *ThreadSafeMap[K, V]) PutIfAbsent(key K, val V) V {
	t.lock.Lock()
	defer t.lock.Unlock()
	v, ok := t.raw_map[key]
	if !ok {
		t.raw_map[key] = val
	}
	return v
}

func (t *ThreadSafeMap[K, V]) Get(key K) (V, bool) {
	t.lock.Lock()
	defer t.lock.Unlock()
	val, ok := t.raw_map[key]
	return val, ok
}

func (t *ThreadSafeMap[K, V]) Delete(key K) {
	t.lock.Lock()
	defer t.lock.Unlock()
	delete(t.raw_map, key)
}

// Clear 清空
func (t *ThreadSafeMap[K, V]) Clear() {
	t.lock.Lock()
	defer t.lock.Unlock()
	for k, _ := range t.raw_map {
		delete(t.raw_map, k)
	}
}

func (t *ThreadSafeMap[K, V]) Keys() []K {
	t.lock.Lock()
	defer t.lock.Unlock()
	ret := make([]K, 0, len(t.raw_map))
	for key, _ := range t.raw_map {
		ret = append(ret, key)
	}
	return ret
}

func (t *ThreadSafeMap[K, V]) Values() []V {
	t.lock.Lock()
	defer t.lock.Unlock()
	ret := make([]V, 0, len(t.raw_map))
	for _, val := range t.raw_map {
		ret = append(ret, val)
	}
	return ret
}

func (t *ThreadSafeMap[K, V]) Range(f func(key K, val V) bool) {
	t.lock.Lock()
	defer t.lock.Unlock()
	for k, v := range t.raw_map {
		if !f(k, v) {
			return
		}
	}
}

func (t *ThreadSafeMap[K, V]) Len() int {
	return len(t.raw_map)
}
