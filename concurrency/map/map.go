package _map

type Map[K comparable, V any] interface {
	Put(key K, val V) V
	PutAll(m map[K]V)
	PutIfAbsent(key K, val V) V
	Get(key K) (V, bool)
	Delete(key K)
	Clear()
	Keys() []K
	Values() []V
	Range(f func(key K, val V) bool)
	Len() int
}
