package rbtree

import (
	"fmt"
	"github.com/koleter/go-util/compare"
	"github.com/koleter/go-util/util"
)

type RBTree[K compare.Comparator[K], V any] struct {
	root *Node[K, V]
	size int
	zero V
}

const (
	red   = false
	black = true
)

type Node[K compare.Comparator[K], V any] struct {
	left   *Node[K, V]
	right  *Node[K, V]
	parent *Node[K, V]
	color  bool
	key    K
	value  V
}

func newNode[K compare.Comparator[K], V any](key K, value V, parent *Node[K, V]) *Node[K, V] {
	return &Node[K, V]{
		left:   nil,
		right:  nil,
		parent: parent,
		color:  red,
		key:    key,
		value:  value,
	}
}

func (node *Node[K, V]) Next() *Node[K, V] {
	if node == nil {
		return nil
	}
	if node.right != nil {
		right := node.right
		for right.left != nil {
			right = right.left
		}
		return right
	}
	p := node.parent
	for p != nil && p.left != node {
		node = p
		p = p.parent
	}
	return p
}

func (node *Node[K, V]) Prev() *Node[K, V] {
	if node == nil {
		return nil
	}
	if node.left != nil {
		left := node.left
		for left.right != nil {
			left = left.right
		}
		return left
	}
	p := node.parent
	for p != nil && p.right != node {
		node = p
		p = p.parent
	}
	return p
}

func colorOf[K compare.Comparator[K], V any](node *Node[K, V]) bool {
	if node == nil {
		return black
	}
	return node.color
}

func (t *RBTree[K, V]) Len() int {
	return t.size
}

func (t *RBTree[K, V]) leftRotate(node *Node[K, V]) {
	if node == nil {
		return
	}
	right := node.right
	right.parent = node.parent
	if node.parent == nil {
		t.root = right
	} else if node.parent.left == node {
		node.parent.left = right
	} else {
		node.parent.right = right
	}

	node.parent = right
	node.right = right.left
	if right.left != nil {
		right.left.parent = node
	}
	right.left = node
}

func (t *RBTree[K, V]) rightRotate(node *Node[K, V]) {
	if node == nil {
		return
	}
	left := node.left
	left.parent = node.parent
	if node.parent == nil {
		t.root = left
	} else if node.parent.left == node {
		node.parent.left = left
	} else {
		node.parent.right = left
	}
	node.parent = left
	node.left = left.right
	if left.right != nil {
		left.right.parent = node
	}
	left.right = node
}

func (t *RBTree[K, V]) afterInsert(node *Node[K, V]) {
	for node != t.root && colorOf(node.parent) == red { //父节点为红色,表示父节点只有一个子节点为自己
		parent := node.parent
		if parent == parent.parent.left { //父节点为祖父的左子节点,祖父节点必为黑色
			uncle := parent.parent.right //获取叔叔节点
			if colorOf(uncle) == red {
				parent.color, uncle.color = black, black
				parent.parent.color = red
				node = parent.parent
			} else { //叔叔节点为黑色,此时叔叔节点只可能为哨兵节点
				if node == parent.right {
					t.leftRotate(parent)
					node, parent = parent, node
				}
				parent.parent.color = red
				parent.color = black
				t.rightRotate(parent.parent)
				break
			}
		} else { //父节点为祖父的右子节点
			uncle := parent.parent.left //获取叔叔节点
			if colorOf(uncle) == red {
				uncle.color = black
				parent.color = black
				parent.parent.color = red
				node = parent.parent
			} else {
				if node == parent.left {
					t.rightRotate(parent)
					node, parent = parent, node
				}
				parent.color = black
				parent.parent.color = red
				t.leftRotate(parent.parent)
				break
			}
		}
	}
	t.root.color = black
}

func (t *RBTree[K, V]) successor(node *Node[K, V]) *Node[K, V] {
	if node == nil {
		return nil
	}
	if node.right == nil {
		return node
	}
	right := node.right
	for right.left != nil {
		right = right.left
	}
	return right
}

func (t *RBTree[K, V]) Insert(key K, value V) *Node[K, V] {
	if util.IsNil(key) {
		panic("key is a null pointer")
	}
	if t.root == nil {
		t.root = &Node[K, V]{
			left:   nil,
			right:  nil,
			parent: nil,
			color:  black,
			key:    key,
			value:  value,
		}
		t.size = 1
		return t.root
	}
	node := t.root
	var parent *Node[K, V]
	for node != nil {
		parent = node
		cmp := node.key.Compare(key)
		if cmp < 0 {
			node = node.right
		} else if cmp == 0 {
			node.value = value
			return node
		} else {
			node = node.left
		}
	}
	n := newNode(key, value, parent)
	if parent.key.Compare(key) > 0 {
		parent.left = n
	} else {
		parent.right = n
	}
	t.afterInsert(n)
	t.size++
	return n
}

/*
从红黑树中删除指定的节点
*/
func (t *RBTree[K, V]) deleteNode(node *Node[K, V]) {
	t.size--
	temp := t.successor(node) //找到后继节点
	node.key = temp.key
	node.value = temp.value

	if temp == t.root { //如果被删除的是根节点
		if t.size == 1 { // 树中有两个节点,表示根节点还有个左孩子,将左孩子设置为根
			t.root = t.root.left
			t.root.parent = nil
			t.root.color = black
		}
		t.root = nil
	} else if temp.right != nil { //后继节点存在右子节点
		temp.right.color = black
		if temp == temp.parent.left {
			temp.parent.left = temp.right
		} else {
			temp.parent.right = temp.right
		}
		temp.right.parent = temp.parent
		temp.right = nil
		temp.parent = nil
	} else { //后继节点无子节点
		if colorOf(temp) == black {
			t.afterDelete(temp)
		}

		var n *Node[K, V]
		if temp.left != nil {
			temp.left.parent = temp.parent
			n = temp.left
		} else if temp.right != nil {
			temp.right.parent = temp.parent
			n = temp.right
		}

		if temp.parent != nil {
			if temp == temp.parent.left {
				temp.parent.left = n
			} else {
				temp.parent.right = n
			}
		}
	}
}

/*
通过key找到红黑树中对应的节点
*/
func (t *RBTree[K, V]) findNodeByKey(key K) *Node[K, V] {
	node := t.root
	for node != nil {
		cmp := node.key.Compare(key)
		if cmp < 0 {
			node = node.right
		} else if cmp == 0 {
			break
		} else {
			node = node.left
		}
	}
	return node
}

/*
删除红黑树中key节点
返回值为true时表示删除成功,为false表示没有这个key的节点
*/
func (t *RBTree[K, V]) Delete(key K) bool {
	node := t.findNodeByKey(key)
	if node == nil {
		return false
	}
	t.deleteNode(node)
	return true
}

func (t *RBTree[K, V]) afterDelete(node *Node[K, V]) {
	for node != t.root && colorOf(node) == black {
		if node == node.parent.left { //是父节点的左节点
			sib := node.parent.right //获取兄弟节点
			if colorOf(sib) == red { //兄弟节点为红色,父节点必为黑色,兄弟节点必有两个黑色子节点,要注意兄弟节点可能有红色的孙子节点
				sib.color = black
				sib.parent.color = red
				t.leftRotate(node.parent)
				sib = node.parent.right
			}

			if colorOf(sib.left) == black && colorOf(sib.right) == black {
				sib.color = red
				if colorOf(node.parent) == red {
					node.parent.color = black
					return
				}
				node = node.parent
			} else { //兄弟节点某个子节点为红色
				if colorOf(sib.left) == red { //如果兄弟节点左节点为红色
					sib.left.color = black
					sib.color = red
					t.rightRotate(sib)
					sib = node.parent.right
				}
				//这下只剩下了兄弟节点为黑色,兄弟节点的右子节点为红色的情况了
				sib.color = colorOf(node.parent)
				node.parent.color = black
				sib.right.color = black
				t.leftRotate(node.parent)
				break
			}
		} else { //是父节点的右节点
			sib := node.parent.left
			if colorOf(sib) == red { //兄弟节点为红色,父节点必为黑色,兄弟节点必有两个黑色子节点,要注意兄弟节点可能有红色的孙子节点
				sib.parent.color = red
				sib.color = black
				t.rightRotate(node.parent)
				sib = node.parent.left
			}

			if colorOf(sib.left) == black && colorOf(sib.right) == black { //左右均为null
				sib.color = red
				if colorOf(node.parent) == red {
					node.parent.color = black
					return
				}
				node = node.parent
			} else {
				if colorOf(sib.right) == red {
					sib.right.color = black
					sib.color = red
					t.leftRotate(sib)
					sib = node.parent.left
				}
				sib.color = node.parent.color
				node.parent.color = black
				sib.left.color = black
				t.rightRotate(node.parent)
				break
			}
		}
	}
	t.root.color = black
}

/*
通过key找到对应的value
*/
func (t *RBTree[K, V]) Get(key K) (V, bool) {
	node := t.findNodeByKey(key)
	if node == nil {
		return t.zero, false
	}
	return node.value, true
}

/*
用于检查红黑树的结构是否正确
msg: 出错时额外打印的信息
*/
func (t *RBTree[K, V]) check(msg interface{}) {
	err := func(errorstr interface{}) {
		fmt.Println(msg)
		panic(errorstr)
	}

	siz := t.size
	var dfs func(node *Node[K, V])
	dfs = func(node *Node[K, V]) {
		if node == nil {
			return
		}
		siz--
		if colorOf(node) == red && (colorOf(node.left) == red || colorOf(node.right) == red) {
			err("出现父节点与子节点都为红色的情况了")
		}
		if node.left != nil {
			if node.left.parent != node {
				err("父子节点关系对应错误")
			}
			if node.key.Compare(node.left.key) <= 0 {
				err(node)
			}
			dfs(node.left)
		}
		if node.right != nil {
			if node.right.parent != node {
				err("父子节点关系对应错误")
			}
			if node.key.Compare(node.right.key) >= 0 {
				err(node)
			}
			dfs(node.right)
		}
	}
	dfs(t.root)
	if siz != 0 {
		err("红黑树中的节点数量不正确")
	}
}

func (t *RBTree[K, V]) lowerNode(key K) *Node[K, V] {
	node := t.root
	for node != nil {
		cmp := key.Compare(node.key)
		// 节点的key小于当前的key,我们要找到小于等于key中最大的那个节点,所以尝试找right
		if cmp > 0 {
			if node.right != nil {
				node = node.right
			} else {
				return node
			}
		} else if cmp < 0 {
			// 节点的key大于当前的key,我们要找到小于等于key中最大的那个节点,所以尝试找比较小的节点
			if node.left != nil {
				node = node.left
			} else {
				// 没有left,通过左分支向上找比较小的父节点就是答案
				parent := node.parent
				// 这里之所以要判断是left分支,是因为确保是从左分支下来的,这样的话父节点必定也是小于key的
				for parent != nil && parent.left == node {
					node = parent
					parent = parent.parent
				}
				return parent
			}

		} else {
			return node.Prev()
		}
	}
	return nil
}

// Lower find the largest value that is smaller than key
func (t *RBTree[K, V]) Lower(key K) (V, bool) {
	node := t.lowerNode(key)
	if node == nil {
		return t.zero, false
	}
	return node.value, true
}

func (t *RBTree[K, V]) higherNode(key K) *Node[K, V] {
	node := t.root
	for node != nil {
		cmp := key.Compare(node.key)
		// 节点的key大于当前的key,我们要找到大于等于key中最大的那个节点,所以尝试找left
		if cmp < 0 {
			if node.left != nil {
				node = node.left
			} else {
				return node
			}
		} else if cmp > 0 {
			if node.right != nil {
				node = node.right
			} else {
				parent := node.parent
				for parent != nil && parent.right == node {
					node = parent
					parent = parent.parent
				}
				return parent
			}

		} else {
			return node.Next()
		}
	}
	return nil
}

func (t *RBTree[K, V]) Higher(key K) (V, bool) {
	node := t.higherNode(key)
	if node == nil {
		return t.zero, false
	}
	return node.value, true
}
