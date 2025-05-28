package queue

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewCircularQueue(t *testing.T) {
	q := NewCoveredCircularQueue[int](5)
	if q.capacity != 5 || q.size != 0 || q.front != 0 || q.rear != 0 {
		t.Errorf("NewCoveredCircularQueue 初始化失败: %+v", q)
	}
}

func TestEnqueue(t *testing.T) {
	q := NewCoveredCircularQueue[int](3)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if q.size != 3 || !q.IsFull() {
		t.Errorf("Enqueue: 队列未满")
	}

	q.Enqueue(4) // 应该覆盖第一个元素
	expected := []int{2, 3, 4}
	var result []int
	q.Range(func(item int) bool {
		result = append(result, item)
		return true
	})
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Enqueue: 覆盖逻辑错误，期望 %v，实际 %v", expected, result)
	}
}

func TestDequeue(t *testing.T) {
	q := NewCoveredCircularQueue[string](3)
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")

	item, ok := q.Dequeue()
	if !ok || item != "a" {
		t.Errorf("Dequeue: 第一次出队错误，期望 'a'")
	}

	item, ok = q.Dequeue()
	if !ok || item != "b" {
		t.Errorf("Dequeue: 第二次出队错误，期望 'b'")
	}

	item, ok = q.Dequeue()
	if !ok || item != "c" {
		t.Errorf("Dequeue: 第三次出队错误，期望 'c'")
	}

	item, ok = q.Dequeue()
	if ok {
		t.Errorf("Dequeue: 空队列不应返回有效值")
	}
}

func TestIsEmpty(t *testing.T) {
	q := NewCoveredCircularQueue[int](2)
	if !q.IsEmpty() {
		t.Errorf("IsEmpty: 初始化应为空")
	}

	q.Enqueue(1)
	if q.IsEmpty() {
		t.Errorf("IsEmpty: 入队后不应为空")
	}

	q.Dequeue()
	if !q.IsEmpty() {
		t.Errorf("IsEmpty: 出队后应为空")
	}
}

func TestIsFull(t *testing.T) {
	q := NewCoveredCircularQueue[int](2)
	if q.IsFull() {
		t.Errorf("IsFull: 初始化不应为满")
	}

	q.Enqueue(1)
	if q.IsFull() {
		t.Errorf("IsFull: 只有一个元素不应为满")
	}

	q.Enqueue(2)
	if !q.IsFull() {
		t.Errorf("IsFull: 两个元素应为满")
	}
}

func TestSizeAndCapacity(t *testing.T) {
	q := NewCoveredCircularQueue[int](5)
	if q.Capacity() != 5 {
		t.Errorf("Capacity: 期望 5")
	}
	if q.Size() != 0 {
		t.Errorf("Size: 初始化应为 0")
	}

	q.Enqueue(1)
	if q.Size() != 1 {
		t.Errorf("Size: 入队后应为 1")
	}

	q.Dequeue()
	if q.Size() != 0 {
		t.Errorf("Size: 出队后应为 0")
	}
}

func TestRange(t *testing.T) {
	q := NewCoveredCircularQueue[int](3)
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	var res []int
	q.Range(func(item int) bool {
		res = append(res, item)
		return true
	})

	expected := []int{10, 20, 30}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Range: 遍历顺序错误，期望 %v，实际 %v", expected, res)
	}
}

func TestRemove(t *testing.T) {
	q := NewCoveredCircularQueue[int](3)
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	q.Range(func(item int) bool {
		if item < 15 {
			q.Dequeue()
			return false
		}
		return true
	})
	assert.Equal(t, 2, q.Size())
	n, _ := q.Dequeue()
	assert.Equal(t, 20, n)
	n, _ = q.Dequeue()
	assert.Equal(t, 30, n)
}

// TestFront_EmptyQueue tests Front() on an empty queue.
func TestFront_EmptyQueue(t *testing.T) {
	q := &CoveredCircularQueue[int]{
		data:     make([]int, 5),
		front:    0,
		rear:     -1,
		size:     0,
		capacity: 5,
	}

	val, ok := q.Front()
	if ok {
		t.Errorf("expected ok to be false for empty queue")
	}
	var zero int
	if val != zero {
		t.Errorf("expected value to be zero value for empty queue, got %v", val)
	}
}

// TestFront_NonEmptyQueue tests Front() on a non-empty queue.
func TestFront_NonEmptyQueue(t *testing.T) {
	q := &CoveredCircularQueue[string]{
		data:     []string{"a", "b", "c"},
		front:    0,
		rear:     2,
		size:     3,
		capacity: 3,
	}

	expected := "a"
	val, ok := q.Front()
	if !ok {
		t.Errorf("expected ok to be true for non-empty queue")
	}
	if val != expected {
		t.Errorf("expected front value %v, got %v", expected, val)
	}
}

// TestBack_EmptyQueue tests Back() on an empty queue.
func TestBack_EmptyQueue(t *testing.T) {
	q := &CoveredCircularQueue[float64]{
		data:     make([]float64, 5),
		front:    0,
		rear:     -1,
		size:     0,
		capacity: 5,
	}

	val, ok := q.Back()
	if ok {
		t.Errorf("expected ok to be false for empty queue")
	}
	var zero float64
	if val != zero {
		t.Errorf("expected value to be zero value for empty queue, got %v", val)
	}
}

// TestBack_NonEmptyQueue tests Back() on a non-empty queue.
func TestBack_NonEmptyQueue(t *testing.T) {
	q := &CoveredCircularQueue[int]{
		data:     []int{10, 20, 30},
		front:    1,
		rear:     2,
		size:     2,
		capacity: 3,
	}

	expected := 30
	val, ok := q.Back()
	if !ok {
		t.Errorf("expected ok to be true for non-empty queue")
	}
	if val != expected {
		t.Errorf("expected back value %v, got %v", expected, val)
	}
}

// TestReverseRange_EmptyQueue 测试空队列的情况
func TestReverseRange_EmptyQueue(t *testing.T) {
	q := &CoveredCircularQueue[int]{
		data:     make([]int, 4),
		front:    0,
		size:     0,
		capacity: 4,
	}

	called := false
	q.ReverseRange(func(item int) bool {
		called = true
		return true
	})

	if called {
		t.Error("Expected fn not to be called for empty queue")
	}
}

// TestReverseRange_AllTrue 测试所有元素都返回 true 的情况
func TestReverseRange_AllTrue(t *testing.T) {
	q := &CoveredCircularQueue[int]{
		data:     make([]int, 4),
		front:    1,
		size:     3,
		capacity: 4,
	}
	// 队列内容：索引1 -> 10, 索引2 -> 20, 索引3 -> 30
	q.data[1] = 10
	q.data[2] = 20
	q.data[3] = 30

	var visited []int
	q.ReverseRange(func(item int) bool {
		visited = append(visited, item)
		return true
	})

	expected := []int{30, 20, 10}
	if !reflect.DeepEqual(visited, expected) {
		t.Errorf("Expected visited %v, got %v", expected, visited)
	}
}

// TestReverseRange_StopInMiddle 测试在中间返回 false 的情况
func TestReverseRange_StopInMiddle(t *testing.T) {
	q := &CoveredCircularQueue[int]{
		data:     make([]int, 5),
		front:    2,
		size:     4,
		capacity: 5,
	}
	// 队列内容：索引2 -> 10, 索引3 -> 20, 索引4 -> 30, 索引0 -> 40
	q.data[2] = 10
	q.data[3] = 20
	q.data[4] = 30
	q.data[0] = 40

	var visited []int
	counter := 0
	q.ReverseRange(func(item int) bool {
		visited = append(visited, item)
		counter++
		return counter < 2 // 第二次调用时返回 false
	})

	expected := []int{40, 30} // 逆序访问应该是 40 -> 30 -> 20 -> 10，但会在第2次中断
	if !reflect.DeepEqual(visited, expected) {
		t.Errorf("Expected visited %v, got %v", expected, visited)
	}
}

func TestCoveredCircularQueue_Clear(t *testing.T) {
	q := NewCoveredCircularQueue[int](2)
	q.Enqueue(3)
	q.Enqueue(1)
	assert.NotEqual(t, 0, q.Size())
	q.Clear()
	assert.Equal(t, 0, q.Size())
}
