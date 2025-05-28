package queue

// CoveredCircularDeque 基于数组的双向循环队列，队列满时插入会覆盖队首或者队尾元素
type CoveredCircularDeque[T any] struct {
	data     []T
	front    int // 队头指针（指向第一个元素）
	rear     int // 队尾指针（指向最后一个元素的下一个位置）
	size     int // 当前元素个数
	capacity int // 总容量
}

// NewCoveredCircularDeque 创建一个新的双向循环队列
func NewCoveredCircularDeque[T any](capacity int) *CoveredCircularDeque[T] {
	return &CoveredCircularDeque[T]{
		data:     make([]T, capacity),
		front:    0,
		rear:     0,
		size:     0,
		capacity: capacity,
	}
}

// IsEmpty 是否为空
func (d *CoveredCircularDeque[T]) IsEmpty() bool {
	return d.size == 0
}

// IsFull 是否已满
func (d *CoveredCircularDeque[T]) IsFull() bool {
	return d.size == d.capacity
}

// Size 获取当前元素数量
func (d *CoveredCircularDeque[T]) Size() int {
	return d.size
}

// Capacity 获取队列总容量
func (d *CoveredCircularDeque[T]) Capacity() int {
	return d.capacity
}

// PushFront 在队头插入元素，若队列满则覆盖最尾部的数据
func (d *CoveredCircularDeque[T]) PushFront(item T) {
	if d.IsFull() {
		// 覆盖：移动 rear 指针（即丢弃最后元素）
		d.rear = (d.rear - 1 + d.capacity) % d.capacity
	} else {
		d.size++
	}
	d.front = (d.front - 1 + d.capacity) % d.capacity
	d.data[d.front] = item
}

// PushBack 在队尾插入元素，若队列满则覆盖最头部的数据
func (d *CoveredCircularDeque[T]) PushBack(item T) {
	if d.IsFull() {
		// 覆盖：移动 front 指针（即丢弃最前元素）
		d.front = (d.front + 1) % d.capacity
	} else {
		d.size++
	}
	d.data[d.rear] = item
	d.rear = (d.rear + 1) % d.capacity
}

// PopFront 删除并返回队头元素
func (d *CoveredCircularDeque[T]) PopFront() (T, bool) {
	var zero T
	if d.IsEmpty() {
		return zero, false
	}
	item := d.data[d.front]
	d.front = (d.front + 1) % d.capacity
	d.size--
	return item, true
}

// PopBack 删除并返回队尾元素
func (d *CoveredCircularDeque[T]) PopBack() (T, bool) {
	var zero T
	if d.IsEmpty() {
		return zero, false
	}
	d.rear = (d.rear - 1 + d.capacity) % d.capacity
	item := d.data[d.rear]
	d.size--
	return item, true
}

// Front 获取队头元素
func (d *CoveredCircularDeque[T]) Front() (T, bool) {
	var zero T
	if d.IsEmpty() {
		return zero, false
	}
	return d.data[d.front], true
}

// Back 获取队尾元素
func (d *CoveredCircularDeque[T]) Back() (T, bool) {
	var zero T
	if d.IsEmpty() {
		return zero, false
	}
	index := (d.rear - 1 + d.capacity) % d.capacity
	return d.data[index], true
}

// Range 遍历所有元素，按从 front 到 rear 的顺序执行函数 fn
func (d *CoveredCircularDeque[T]) Range(fn func(T) bool) {
	if d.IsEmpty() {
		return
	}
	for i := 0; i < d.size; i++ {
		index := (d.front + i) % d.capacity
		if !fn(d.data[index]) {
			return
		}
	}
}

// ReverseRange 遍历队列，按逆序执行函数 fn
func (d *CoveredCircularDeque[T]) ReverseRange(fn func(item T) bool) {
	if d.IsEmpty() {
		return
	}

	for i := d.size - 1; i >= 0; i-- {
		index := (d.front + i) % d.capacity
		if !fn(d.data[index]) {
			return
		}
	}
}

func (q *CoveredCircularDeque[T]) Clear() {
	q.size = 0
	q.rear = q.front
}
