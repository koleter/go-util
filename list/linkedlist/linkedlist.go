package linkedlist

// Node 定义单向链表的节点
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// LinkedList 定义单向链表,FIFO
type LinkedList[T any] struct {
	head, tail *Node[T]
	size       int
}

func (l *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = l.head
	} else {
		l.tail.Next = newNode
		l.tail = newNode
	}
	l.size++
}

func (l *LinkedList[T]) Len() int {
	return l.size
}

func (l *LinkedList[T]) Pop() (T, bool) {
	var zero T
	if l.head == nil {
		return zero, false
	}
	res := l.head.Value
	l.head = l.head.Next
	l.size--
	return res, true
}
