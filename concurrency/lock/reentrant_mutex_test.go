package lock

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestReentrantMutex(t *testing.T) {
	var lock ReentrantMutex

	var wg sync.WaitGroup
	wg.Add(2)

	var c int
	go func() {
		for i := 0; i < 10000; i++ {
			lock.Lock()
			c++
			lock.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			lock.Lock()
			c++
			lock.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
	assert.Equal(t, 20000, c)
}
