package coroutine

import (
	"fmt"
	"github.com/koleter/go-util/list/linkedlist"
	"log"
	"runtime/debug"
	"sync"
)

type RoutinePool struct {
	tasks   linkedlist.LinkedList[func()]
	wg      sync.WaitGroup
	mu      sync.Mutex
	cond    *sync.Cond
	stopped bool
}

// NewRoutinePool 创建一个新的实例
func NewRoutinePool(workerNum int) *RoutinePool {
	if workerNum <= 0 {
		panic(fmt.Sprintf("workerNum is less than or equal to 0, workerNum: %d", workerNum))
	}
	pool := &RoutinePool{
		tasks: linkedlist.LinkedList[func()]{},
	}
	pool.cond = sync.NewCond(&pool.mu)
	for i := 0; i < workerNum; i++ {
		pool.wg.Add(1)
		go pool.worker()
	}
	return pool
}

// Stop 停止协程池不再接收新任务
func (p *RoutinePool) Stop() {
	if p.stopped {
		return
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.stopped = true
	p.cond.Broadcast()
}

// Wait 等待所有任务执行完毕,调用该函数前必须调用 Stop
func (p *RoutinePool) Wait() {
	p.wg.Wait()
}

func runTask(task func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Task panicked: %v\nstack: %s", err, string(debug.Stack()))
		}
	}()
	task()
}

// worker 是每个工作者的主循环
func (p *RoutinePool) worker() {
	defer p.wg.Done()
	for !p.stopped || p.tasks.Len() != 0 {
		p.mu.Lock()
		if p.tasks.Len() == 0 {
			if p.stopped {
				p.mu.Unlock()
				return
			}
			p.cond.Wait()
		}
		task, ok := p.tasks.Pop()
		p.mu.Unlock()
		if !ok {
			continue
		}
		runTask(task)
	}
}

// Submit 提交一个任务
func (p *RoutinePool) Submit(task func()) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if !p.stopped {
		p.tasks.Append(task)
		p.cond.Signal()
	}
}
