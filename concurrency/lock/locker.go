package lock

type Locker interface {
	WithLock(func())
}
