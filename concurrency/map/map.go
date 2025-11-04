package _map

import "github.com/koleter/go-util/concurrency/lock"

type Map[K comparable, V any] interface {
	Put(key K, val V) V
	PutAll(m map[K]V)
	PutIfAbsent(key K, val V) V
	Get(key K) (V, bool)
	Delete(key K) (V, bool)
	Clear()
	Keys() []K
	Values() []V
	Range(f func(key K, val V) bool)
	Len() int
}

type ConcurrentMap[K comparable, V any] interface {
	lock.Locker
	Map[K, V]
}
