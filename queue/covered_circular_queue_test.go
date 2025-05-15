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
