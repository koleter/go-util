package Map

type node[K comparable, V any] struct {
	children map[K]*node[K, V]
	hasVal   bool
	val      V
}

func newNode[K comparable, V any]() *node[K, V] {
	return &node[K, V]{
		children: make(map[K]*node[K, V]),
	}
}

func (n *node[K, V]) getAllValues(res *[]V) {
	if n.hasVal {
		*res = append(*res, n.val)
	}
	for _, node := range n.children {
		node.getAllValues(res)
	}
}

type MultiKeyMap[K comparable, V any] struct {
	root *node[K, V]
}

func NewMultiKeyMap[K comparable, V any]() *MultiKeyMap[K, V] {
	return &MultiKeyMap[K, V]{
		root: newNode[K, V](),
	}
}

// Put 设置键值对，返回旧值
func (m *MultiKeyMap[K, V]) Put(keys []K, val V) V {
	var res V
	node := m.root
	for _, key := range keys {
		next := node.children[key]
		if next == nil {
			next = newNode[K, V]()
			node.children[key] = next
		}
		node = next
	}
	if node.hasVal {
		res = node.val
	} else {
		node.hasVal = true
	}
	node.val = val
	return res
}

func (m *MultiKeyMap[K, V]) Get(keys []K) (V, bool) {
	var zero V
	node := m.root
	for _, key := range keys {
		next := node.children[key]
		if next == nil {
			return zero, false
		}
		node = next
	}
	if node.hasVal {
		return node.val, true
	}
	return zero, false
}

func (m *MultiKeyMap[K, V]) GetPrefix(keys []K) []V {
	var res []V
	node := m.root
	for _, key := range keys {
		next := node.children[key]
		if next == nil {
			return res
		}
		node = next
	}
	node.getAllValues(&res)
	return res
}

func (m *MultiKeyMap[K, V]) delete(node *node[K, V], keys []K) (V, bool) {
	var res V
	if node == nil {
		return res, false
	}
	if len(keys) != 0 {
		n := node.children[keys[0]]
		res, exist := m.delete(n, keys[1:])
		if exist && len(n.children) == 0 {
			delete(node.children, keys[0])
		}
		return res, exist
	}
	if node.hasVal {
		node.hasVal = false
		return node.val, true
	}
	return res, false
}

func (m *MultiKeyMap[K, V]) Delete(keys []K) (V, bool) {
	var zero V
	if len(keys) == 0 {
		return zero, false
	}
	return m.delete(m.root, keys)
}
