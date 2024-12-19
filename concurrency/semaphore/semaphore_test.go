package semaphore_test

import (
	"github.com/koleter/go-util/concurrency/semaphore"
	"testing"
	"time"
)

// TestBasicAcquireRelease 测试基本的 Acquire 和 Release 功能
func TestBasicAcquireRelease(t *testing.T) {
	sem := semaphore.NewSemaphore(2)

	if err := sem.Acquire(1); err != nil {
		t.Errorf("Acquire failed: %v", err)
	}

	sem.Release(1)
}

// TestCloseSemaphore 测试关闭信号量后的 Acquire 和 Release 情况
func TestCloseSemaphore(t *testing.T) {
	sem := semaphore.NewSemaphore(2)

	sem.Close()

	if err := sem.Acquire(1); err != nil {
		t.Errorf(err.Error())
	}

	sem.Release(1) // 这个操作应该被忽略
}

// TestConcurrentAcquireRelease 测试并发 Acquire 和 Release 情况
func TestConcurrentAcquireRelease(t *testing.T) {
	sem := semaphore.NewSemaphore(2)
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func(id int) {
			if err := sem.Acquire(1); err != nil {
				t.Errorf("Acquire failed for goroutine %d: %v", id, err)
				done <- false
				return
			}
			time.Sleep(100 * time.Millisecond)
			sem.Release(1)
			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		if !<-done {
			t.Errorf("Goroutine %d failed", i)
		}
	}
}
