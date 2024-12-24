package coroutine

import (
	"sync"
	"testing"
	"time"
)

func TestNewRoutinePool_PositiveWorkerNum_Success(t *testing.T) {
	pool := NewRoutinePool(3)
	if pool == nil {
		t.Error("Expected a non-nil RoutinePool")
	}
}

func TestNewRoutinePool_ZeroOrNegativeWorkerNum_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected a panic for zero or negative workerNum")
		}
	}()
	NewRoutinePool(0)
}

func TestStop_WithTasks_Completes(t *testing.T) {
	pool := NewRoutinePool(2)
	var wg sync.WaitGroup
	wg.Add(2)
	pool.Submit(func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
	})
	pool.Submit(func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
	})
	pool.Stop()
	wg.Wait()
}

func TestStop_WithoutTasks_Completes(t *testing.T) {
	pool := NewRoutinePool(2)
	pool.Stop()
	pool.Stop()
}

func TestSubmit_WhenStopped_IgnoresTask(t *testing.T) {
	pool := NewRoutinePool(1)
	pool.Stop()
	pool.Submit(func() {
		t.Error("Task should not be executed")
	})
}

func TestWorker_ExecutesTasks(t *testing.T) {
	pool := NewRoutinePool(1)
	var wg sync.WaitGroup
	wg.Add(1)
	pool.Submit(func() {
		defer wg.Done()
	})
	wg.Wait()
	pool.Stop()
}