package lock

import (
	"go-util/g"
	"runtime"
	"sync/atomic"
	"unsafe"
)

// ReentrantMutex 可重入锁
type ReentrantMutex struct {
	owner unsafe.Pointer // 当前协程的指针
	count int32          // 锁的嵌套深度
}

// Lock 尝试获取锁
func (r *ReentrantMutex) Lock() {
	gp := g.G()
	if atomic.LoadPointer(&r.owner) == gp {
		// 如果当前 goroutine 已经拥有锁，则递增计数器
		atomic.AddInt32(&r.count, 1)
		return
	}

	// 否则，阻塞当前 goroutine 直到锁可用
	var iter int
	for !atomic.CompareAndSwapPointer(&r.owner, unsafe.Pointer(nil), gp) {
		iter++
		if iter == 4 {
			iter = 0
			// 自旋4次未获得锁，让出cpu
			runtime.Gosched()
		}
	}
	atomic.StoreInt32(&r.count, 1)
}

// Unlock 释放锁
func (r *ReentrantMutex) Unlock() {
	gp := g.G()
	if atomic.LoadPointer(&r.owner) != gp {
		panic("unlock of unlocked reentrant mutex")
	}
	if atomic.AddInt32(&r.count, -1) == 0 {
		atomic.StorePointer(&r.owner, unsafe.Pointer(nil))
	}
}
