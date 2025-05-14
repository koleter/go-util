package dlinkedlist

type Node[T any] struct {
	Value      T
	prev, next *Node[T]
	link       *DoublyLinkedList[T]
}

func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// DoublyLinkedList 双向链表
type DoublyLinkedList[T any] struct {
	head, tail *Node[T]
	size       int
}

func (l *DoublyLinkedList[T]) Clear() {
	l.size = 0
	l.head = nil
	l.tail = nil
}

func (l *DoublyLinkedList[T]) Head() *Node[T] {
	return l.head
}

func (l *DoublyLinkedList[T]) Tail() *Node[T] {
	return l.tail
}

// PushBack 插入到尾部
func (l *DoublyLinkedList[T]) PushBack(value T) *Node[T] {
	newNode := &Node[T]{
		link:  l,
		Value: value,
	}
	if l.size == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
	return newNode
}

// PushFront 插入到头部
func (l *DoublyLinkedList[T]) PushFront(value T) *Node[T] {
	newNode := &Node[T]{
		link:  l,
		Value: value,
	}
	if l.size == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	}
	l.size++
	return newNode
}

func (l *DoublyLinkedList[T]) Len() int {
	return l.size
}

// PopFront 弹出第一个值
func (l *DoublyLinkedList[T]) PopFront() (T, bool) {
	var zero T
	if l.size == 0 {
		return zero, false
	}
	res := l.head.Value
	if l.size == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.prev = nil
	}
	l.size--
	return res, true
}

// PopBack 弹出最后一个值
func (l *DoublyLinkedList[T]) PopBack() (T, bool) {
	var zero T
	if l.size == 0 {
		return zero, false
	}
	res := l.tail.Value
	if l.size == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}
	l.size--
	return res, true
}

// Remove 将节点从链表中删除
func (l *DoublyLinkedList[T]) Remove(n *Node[T]) {
	if n == nil || n.link != l {
		panic("remove a node don't belong this link")
	}
	defer func() {
		l.size--
		n.link = nil
	}()
	if n == l.head && n == l.tail {
		// 节点既是头又是尾（唯一节点）
		l.head = nil
		l.tail = nil
	} else if n == l.head {
		// 头节点
		l.head = n.next
		l.head.prev = nil
	} else if n == l.tail {
		// 尾节点
		l.tail = n.prev
		l.tail.next = nil
	} else {
		// 中间节点
		prev := n.prev
		next := n.next
		prev.next = next
		next.prev = prev
	}
}
