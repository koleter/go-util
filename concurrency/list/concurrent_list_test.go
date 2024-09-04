package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_threadSafeList_Append(t *testing.T) {
	safeList := NewThreadSafeList([]int{})
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
		return val == 3
	}))
	safeList.Remove(2)
	assert.False(t, safeList.Contain(func(i int, val int) bool {
		return val == 3
	}))
}
