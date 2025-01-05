package compare

type Comparator[T any] interface {
	Compare(other T) int
}
