package syncPool

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type node[T any] struct {
	val T
}

func TestNewSyncPool_same_pool_when_type_is_same(t *testing.T) {
	syncPool1 := NewSyncPool[int]()
	syncPool2 := NewSyncPool[int]()
	assert.Equal(t, syncPool1, syncPool2)
}

func TestNewSyncPool_same_pool_when_Generics_type_is_same(t *testing.T) {
	syncPool1 := NewSyncPool[node[string]]()
	syncPool2 := NewSyncPool[node[string]]()
	assert.Equal(t, syncPool1, syncPool2)
}

func TestNewSyncPool_not_same_pool_when_Generics_type_is_not_same(t *testing.T) {
	syncPool1 := NewSyncPool[node[string]]()
	syncPool2 := NewSyncPool[node[int]]()
	assert.NotEqual(t, syncPool1, syncPool2)
	syncPool3 := NewSyncPool[node[int64]]()
	assert.NotEqual(t, syncPool3, syncPool2)
}

func TestSyncPool_Get_reuse_object(t *testing.T) {
	syncPool := NewSyncPool[int]()
	firstGet := syncPool.Get()
	*firstGet = 4
	first := fmt.Sprintf("%p", firstGet)
	syncPool.Put(firstGet)
	secondGet := syncPool.Get()
	second := fmt.Sprintf("%p", secondGet)
	assert.Equal(t, first, second)
	assert.Equal(t, *firstGet, *secondGet)
}
