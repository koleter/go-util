package syncPool

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSyncPool_same_pool_when_type_is_same(t *testing.T) {
	syncPool1 := NewSyncPool[int]()
	syncPool2 := NewSyncPool[int]()
	assert.Equal(t, syncPool1, syncPool2)
}

func TestSyncPool_Get_reuse_object(t *testing.T) {
	syncPool := NewSyncPool[int]()
	get := syncPool.Get()
	first := fmt.Sprintf("%p", get)
	syncPool.Put(get)
	get = syncPool.Get()
	second := fmt.Sprintf("%p", get)
	assert.Equal(t, first, second)
}
