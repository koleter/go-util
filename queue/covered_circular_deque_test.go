package queue

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// TestNewCircularDeque tests the creation of a new circular deque
func TestNewCircularDeque(t *testing.T) {
	deque := NewCoveredCircularDeque[int](3)
	if deque == nil {
		t.Fatal("Expected non-nil deque")
	}
	if deque.Capacity() != 3 {
		t.Errorf("Expected capacity 3, got %d", deque.Capacity())
	}
	if !deque.IsEmpty() {
		t.Errorf("Expected empty deque")
	}
	if deque.Size() != 0 {
		t.Errorf("Expected size 0, got %d", deque.Size())
	}
}

// TestPushFront_EmptyQueue tests pushing to front of an empty deque
func TestPushFront_EmptyQueue(t *testing.T) {
	deque := NewCoveredCircularDeque[int](2)
	deque.PushFront(1)

	if deque.Size() != 1 {
		t.Errorf("Expected size 1, got %d", deque.Size())
	}

	if val, ok := deque.Front(); !ok || val != 1 {
		t.Errorf("Expected front value 1, got %v", val)
	}

	if val, ok := deque.Back(); !ok || val != 1 {
		t.Errorf("Expected back value 1, got %v", val)
	}
}

// TestPushBack_EmptyQueue tests pushing to back of an empty deque
func TestPushBack_EmptyQueue(t *testing.T) {
	deque := NewCoveredCircularDeque[int](2)
	deque.PushBack(1)

	if deque.Size() != 1 {
		t.Errorf("Expected size 1, got %d", deque.Size())
	}

	if val, ok := deque.Front(); !ok || val != 1 {
		t.Errorf("Expected front value 1, got %v", val)
	}

	if val, ok := deque.Back(); !ok || val != 1 {
		t.Errorf("Expected back value 1, got %v", val)
	}
}

// TestPushFront_Overwrite tests pushing to front when deque is full
func TestPushFront_Overwrite(t *testing.T) {
	deque := NewCoveredCircularDeque[int](2)
	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushFront(3) // Should overwrite 2

	expected := []int{3, 1}
	var result []int
	deque.Range(func(item int) bool {
		result = append(result, item)
		return true
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestPushBack_Overwrite tests pushing to back when deque is full
func TestPushBack_Overwrite(t *testing.T) {
	deque := NewCoveredCircularDeque[int](2)
	deque.PushBack(1)
	deque.PushFront(2)
	deque.PushBack(3) // Should overwrite 2

	expected := []int{1, 3}
	var result []int
	deque.Range(func(item int) bool {
		result = append(result, item)
		return true
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestPopFront tests popping from the front of the deque
func TestPopFront(t *testing.T) {
	deque := NewCoveredCircularDeque[int](3)
	deque.PushBack(1)
	deque.PushBack(2)

	val, ok := deque.PopFront()
	if !ok || val != 1 {
		t.Errorf("Expected popped value 1, got %v", val)
	}

	if size := deque.Size(); size != 1 {
		t.Errorf("Expected size 1, got %d", size)
	}

	val, ok = deque.Front()
	if !ok || val != 2 {
		t.Errorf("Expected new front value 2, got %v", val)
	}
}

// TestPopBack tests popping from the back of the deque
func TestPopBack(t *testing.T) {
	deque := NewCoveredCircularDeque[int](3)
	deque.PushBack(1)
	deque.PushBack(2)

	val, ok := deque.PopBack()
	if !ok || val != 2 {
		t.Errorf("Expected popped value 2, got %v", val)
	}

	if size := deque.Size(); size != 1 {
		t.Errorf("Expected size 1, got %d", size)
	}

	val, ok = deque.Back()
	if !ok || val != 1 {
		t.Errorf("Expected new back value 1, got %v", val)
	}
}

// TestEdgeCases tests edge cases like empty deque operations and wrapping pointers
func TestEdgeCases(t *testing.T) {
	deque := NewCoveredCircularDeque[int](2)

	// Test operations on empty deque
	if _, ok := deque.PopFront(); ok {
		t.Error("Expected false when popping from empty deque")
	}

	if _, ok := deque.PopBack(); ok {
		t.Error("Expected false when popping from empty deque")
	}

	if _, ok := deque.Front(); ok {
		t.Error("Expected false when getting front of empty deque")
	}

	if _, ok := deque.Back(); ok {
		t.Error("Expected false when getting back of empty deque")
	}

	// Test pointer wrapping
	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushBack(3)  // Overwrite oldest element (1)
	deque.PushFront(4) // Overwrite oldest element (3)

	expected := []int{4, 2}
	var result []int
	deque.Range(func(item int) bool {
		result = append(result, item)
		return true
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestRangeEarlyTermination tests that returning false from range stops iteration
func TestRangeEarlyTermination(t *testing.T) {
	deque := NewCoveredCircularDeque[int](3)
	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushBack(3)

	count := 0
	deque.Range(func(item int) bool {
		count++
		return count < 1 // stop after first element
	})

	if count != 1 {
		t.Errorf("Expected only 1 iteration, got %d", count)
	}
}

// TestReverseRange_NormalTraversal 测试正常遍历所有元素
func TestReverseRange_NormalTraversal(t *testing.T) {
	deque := &CoveredCircularDeque[int]{
		data:     []int{10, 20, 30, 40},
		front:    0,
		size:     4,
		capacity: 4,
	}

	expected := []int{40, 30, 20, 10}
	var result []int

	deque.ReverseRange(func(item int) bool {
		result = append(result, item)
		return true
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestReverseRange_StopAtSecondElement 测试在第二个元素停止
func TestReverseRange_StopAtSecondElement(t *testing.T) {
	deque := &CoveredCircularDeque[int]{
		data:     []int{10, 20, 30, 40},
		front:    0,
		size:     4,
		capacity: 4,
	}

	expected := []int{40, 30}
	var result []int

	deque.ReverseRange(func(item int) bool {
		result = append(result, item)
		return len(result) < 2 // stop after second call
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestReverseRange_SingleElement 测试单个元素情况
func TestReverseRange_SingleElement(t *testing.T) {
	deque := &CoveredCircularDeque[int]{
		data:     []int{99},
		front:    0,
		size:     1,
		capacity: 1,
	}

	expected := []int{99}
	var result []int

	deque.ReverseRange(func(item int) bool {
		result = append(result, item)
		return true
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestReverseRange_CircularBuffer 测试 front 不为零的情况
func TestReverseRange_CircularBuffer(t *testing.T) {
	deque := &CoveredCircularDeque[int]{
		data:     []int{30, 40, 10, 20}, // 实际逻辑上 [10, 20, 30, 40]
		front:    2,                     // front points to index 2
		size:     4,
		capacity: 4,
	}

	expected := []int{40, 30, 20, 10}
	var result []int

	deque.ReverseRange(func(item int) bool {
		result = append(result, item)
		return true
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCoveredCircularDeque_Clear(t *testing.T) {
	q := NewCoveredCircularDeque[int](3)
	q.PushBack(3)
	q.PushBack(1)
	assert.NotEqual(t, 0, q.Size())
	q.Clear()
	assert.Equal(t, 0, q.Size())
}
