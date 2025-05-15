package queue

// CoveredCircularQueue 循环队列结构体,队列满会覆盖最先插入的数据
type CoveredCircularQueue[T any] struct {
	data     []T
	front    int
	rear     int
	size     int // 当前元素数量
	capacity int // 总容量
}

// NewCircularQueue 创建一个新的循环队列
func NewCircularQueue[T any](capacity int) *CoveredCircularQueue[T] {
	return &CoveredCircularQueue[T]{
		data:     make([]T, capacity),
		front:    0,
		rear:     0,
		size:     0,
		capacity: capacity,
	}
}

// Enqueue 入队
func (q *CoveredCircularQueue[T]) Enqueue(item T) {
	if q.size == q.capacity {
		// 队列已满，覆盖最旧元素
		q.front = (q.front + 1) % q.capacity
	} else {
		q.size++
	}
	q.data[q.rear] = item
	q.rear = (q.rear + 1) % q.capacity
}

// Dequeue 出队
func (q *CoveredCircularQueue[T]) Dequeue() (T, bool) {
	var zero T
	if q.IsEmpty() {
		return zero, false
	}
	item := q.data[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--
	return item, true
}

// IsEmpty 是否为空
func (q *CoveredCircularQueue[T]) IsEmpty() bool {
	return q.size == 0
}

// IsFull 是否已满
func (q *CoveredCircularQueue[T]) IsFull() bool {
	return q.size == q.capacity
}

// Size 获取当前元素数量
func (q *CoveredCircularQueue[T]) Size() int {
	return q.size
}

// Capacity 获取队列总容量
func (q *CoveredCircularQueue[T]) Capacity() int {
	return q.capacity
}

// Range 遍历队列，按顺序执行函数 fn
func (q *CoveredCircularQueue[T]) Range(fn func(item T) bool) {
	if q.IsEmpty() {
		return
	}

	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		if !fn(q.data[index]) {
			return
		}
	}
}
