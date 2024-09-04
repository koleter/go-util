package list

type List[T any] interface {
	Append(element ...T)
	Get(i int) T
	Len() int
	Range(f func(int, T) bool)
	Contain(f func(int, T) bool) bool
	Filter(f func(int, T) bool) []T
	Clear()
}
