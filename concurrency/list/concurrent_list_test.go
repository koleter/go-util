package list

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func Test_threadSafeList_Append(t *testing.T) {
	var safeList ConcurrentList[int] = NewThreadSafeList([]int{})
	total := 3
	for i := 0; i < total; i++ {
		safeList.Append(i)
	}
	assert.Equal(t, total, safeList.Len())
	for i := 0; i < total; i++ {
		assert.Equal(t, i, safeList.Get(i))
	}
}

func Test_threadSafeList_Clear(t *testing.T) {
	safeList := NewThreadSafeList([]int{})
	total := 3
	for i := 0; i < total; i++ {
		safeList.Append(i)
	}
	assert.Equal(t, total, safeList.Len())
	safeList.Clear()
	assert.Equal(t, 0, safeList.Len())
}

func Test_threadSafeList_Contain(t *testing.T) {
	safeList := NewThreadSafeList([]int{})
	safeList.Append(2)
	safeList.Append(4)
	assert.True(t, safeList.Contain(func(i int, val int) bool {
		return val == 4
	}))

	assert.True(t, safeList.Contain(func(i int, val int) bool {
		return val == 2
	}))

	assert.False(t, safeList.Contain(func(i int, val int) bool {
		return val == 9
	}))
}

func Test_threadSafeList_Filter(t *testing.T) {
	safeList := NewThreadSafeList([]int{})
	total := 10
	for i := 0; i < total; i++ {
		safeList.Append(i)
	}
	filter := safeList.Filter(func(i int, val int) bool {
		return val&1 == 0
	})
	assert.Equal(t, []int{0, 2, 4, 6, 8}, filter)
}

func Test_threadSafeList_Remove(t *testing.T) {
	safeList := NewThreadSafeList([]int{})
	for i := 0; i < 5; i++ {
		safeList.Append(i)
	}
	assert.True(t, safeList.Contain(func(i int, val int) bool {
		return val == 2
	}))
	safeList.Remove(2)
	assert.False(t, safeList.Contain(func(i int, val int) bool {
		return val == 2
	}))
}

func TestThreadSafeList_Sort(t *testing.T) {
	safeList := NewThreadSafeList([]int{})
	total := 10
	for i := 0; i < total; i++ {
		safeList.Append(rand.Int())
	}
	safeList.Sort(func(i, j int) bool {
		i1 := safeList.Get(i)
		i2 := safeList.Get(j)
		return i1 < i2
	})
	for i := 1; i < total; i++ {
		assert.True(t, safeList.Get(i-1) < safeList.Get(i))
	}
}

func TestThreadSafeList_RemoveFunc(t *testing.T) {
	safeList := NewThreadSafeList([]int{})
	total := 10
	for i := 0; i < total; i++ {
		safeList.Append(i)
	}
	assert.True(t, safeList.Contain(func(i int, n int) bool {
		return n == 4
	}))
	var cnt int
	safeList.RemoveFunc(func(n int) bool {
		cnt++
		return n == 4
	})
	// Make sure the number of cycles is correct
	assert.Equal(t, total, cnt)
	assert.Equal(t, total-1, safeList.Len())
	assert.False(t, safeList.Contain(func(i int, n int) bool {
		return n == 4
	}))
}

func TestThreadSafeList_Set(t *testing.T) {
	safeList := NewThreadSafeList([]int{})
	safeList.Append(5)
	assert.Equal(t, 5, safeList.Get(0))
	assert.Equal(t, 1, safeList.Len())
	oldVal := safeList.Set(0, 9)
	assert.Equal(t, 5, oldVal)
	assert.Equal(t, 9, safeList.Get(0))
}

func TestThreadSafeList_WithLock(t *testing.T) {
	safeList := NewThreadSafeList([]int{})
	safeList.Append(1)
	loop := true
	go func() {
		for loop {
			safeList.WithLock(func() {
				a := 1
				safeList.Set(0, a)
				get := safeList.Get(0)
				assert.Equal(t, a, get)
			})
		}
	}()

	go func() {
		for loop {
			safeList.WithLock(func() {
				a := 2
				safeList.Set(0, a)
				get := safeList.Get(0)
				assert.Equal(t, a, get)
			})
		}
	}()

	time.Sleep(5 * time.Second)
	loop = false
}
