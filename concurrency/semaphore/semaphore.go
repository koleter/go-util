package semaphore

import (
	"fmt"
	"sync"
)

type Semaphore struct {
	mu     sync.Mutex
	cond   *sync.Cond
	count  int
	max    int
	closed bool
}

// NewSemaphore 创建一个新的信号量实例
func NewSemaphore(max int) *Semaphore {
	s := &Semaphore{
		max: max,
	}
	s.cond = sync.NewCond(&s.mu)
	return s
}

// Acquire 请求信号量
func (s *Semaphore) Acquire(cnt int) error {
	if cnt <= 0 {
		panic("cnt <= 0")
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	for s.count+cnt > s.max {
		if s.closed {
			return fmt.Errorf("semaphore is closed")
		}
		s.cond.Wait()
	}

	s.count += cnt
	return nil
}

// Release 释放信号量
func (s *Semaphore) Release(cnt int) {
	if cnt <= 0 {
		panic("cnt <= 0")
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.count < cnt {
		fmt.Errorf("release without acquire")
	}

	s.count -= cnt
	s.cond.Signal()
}

// Close 关闭信号量
func (s *Semaphore) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.closed = true
	s.cond.Broadcast()
}
