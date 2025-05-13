package syncPool

import (
	"reflect"
	"sync"
)

// SyncPool 对sync.Pool进行的封装, 使用如果是相同类型, 会使用相同的sync.Pool
type SyncPool[T any] struct {
	pool sync.Pool
}

var poolMap = make(map[reflect.Type]any)
var poolMapLock sync.RWMutex

func NewSyncPool[T any]() *SyncPool[T] {
	var t T
	typeOf := reflect.TypeOf(t)
	if typeOf == nil || typeOf.Kind() == reflect.Interface {
		panic("cannot create SyncPool for interface{} or nil type")
	}

	if typeOf.Kind() == reflect.Ptr {
		panic("SyncPool does not support pointer types")
	}

	poolMapLock.RLock()
	sp, ok := poolMap[typeOf]
	poolMapLock.RUnlock()
	if ok {
		return sp.(*SyncPool[T])
	}
	poolMapLock.Lock()
	defer poolMapLock.Unlock()
	sp, ok = poolMap[typeOf]
	if ok {
		return sp.(*SyncPool[T])
	}
	s := &SyncPool[T]{
		pool: sync.Pool{
			New: func() any {
				return new(T)
			},
		},
	}
	poolMap[typeOf] = s
	return s
}

func (s *SyncPool[T]) Get() *T {
	return s.pool.Get().(*T)
}

func (s *SyncPool[T]) Put(t *T) {
	s.pool.Put(t)
}
