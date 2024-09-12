package _map

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"sync"
	"testing"
	"time"
)

func TestThreadSafeMap_concurrent_safe(t *testing.T) {
	var safeMap ConcurrentMap[int, int] = NewThreadSafeMap(map[int]int{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 10000; i++ {
			safeMap.Put(i, i)
		}
		wg.Done()
	}()

	go func() {
		for i := 10000; i < 20000; i++ {
			safeMap.Put(i, i)
		}
		wg.Done()
	}()

	wg.Wait()
	assert.Equal(t, 20000, safeMap.Len())
}

func TestThreadSafeMap_Reentrant(t *testing.T) {
	safeMap := NewThreadSafeMap(map[int]int{})
	for i := 0; i < 5; i++ {
		safeMap.Put(i, i)
	}
	safeMap.Range(func(key int, val int) bool {
		if key == 3 {
			safeMap.Delete(key)
		}
		return true
	})
	assert.Equal(t, 4, safeMap.Len())
	_, b := safeMap.Get(3)
	assert.False(t, b)
}

// ThreadSafeMap在遍历时进行删除操作确保仍可以遍历所有的元素
func TestThreadSafeMap_Delete_when_Range(t *testing.T) {
	safeMap := NewThreadSafeMap(map[int]int{})
	total := 100
	for i := 0; i < total; i++ {
		safeMap.Put(i, i)
	}
	var visited []int
	safeMap.Range(func(key int, val int) bool {
		if key&1 == 0 {
			visited = append(visited, key)
		} else {
			safeMap.Delete(key)
		}
		return true
	})
	sort.Ints(visited)
	var expect []int
	for i := 0; i < total; i++ {
		if i&1 == 0 {
			expect = append(expect, i)
		}
	}
	assert.Equal(t, total/2, len(visited))
	assert.Equal(t, expect, visited)
}

func TestThreadSafeMap_WithLock(t *testing.T) {
	safeMap := NewThreadSafeMap(map[int]int{})
	loop := true
	go func() {
		for loop {
			safeMap.WithLock(func() {
				safeMap.Put(1, 2)
				get, _ := safeMap.Get(1)
				assert.Equal(t, 2, get)
			})
		}
	}()

	go func() {
		for loop {
			safeMap.WithLock(func() {
				safeMap.Put(1, 3)
				get, _ := safeMap.Get(1)
				assert.Equal(t, 3, get)
			})
		}
	}()

	time.Sleep(5 * time.Second)
	loop = false
}
