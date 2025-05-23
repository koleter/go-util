package rbtree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 比较器函数
func intCompare(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func TestRBTree_Delete(t *testing.T) {
	tree := NewRBTree[int, int](intCompare)
	tree.Insert(3, 5)
	get, exist := tree.Get(3)
	assert.True(t, exist)
	assert.Equal(t, 5, get)

	tree.Delete(3)
	_, exist = tree.Get(3)
	assert.False(t, exist)
}

func TestRBTree_Higher(t *testing.T) {
	tree := NewRBTree[int, int](intCompare)
	for i := 0; i < 20; i++ {
		tree.Insert(i, i)
		assert.Equal(t, i+1, tree.Len())
	}

	higher, exist := tree.Higher(5)
	assert.True(t, exist)
	assert.Equal(t, 6, higher)
}

func TestRBTree_Lower(t *testing.T) {
	tree := NewRBTree[int, int](intCompare)
	for i := 0; i < 20; i++ {
		tree.Insert(i, i)
		assert.Equal(t, i+1, tree.Len())
	}

	lower, exist := tree.Lower(13)
	assert.True(t, exist)
	assert.Equal(t, 12, lower)
}

func TestRBTree_Len(t *testing.T) {
	tree := NewRBTree[int, int](intCompare)
	for i := 0; i < 20; i++ {
		tree.Insert(i, i)
		assert.Equal(t, i+1, tree.Len())
	}

	// insert a exist Key, length should not change
	tree.Insert(2, 2)
	assert.Equal(t, 20, tree.Len())
}

func TestNewRBTree(t *testing.T) {
	assert.Panics(t, func() {
		NewRBTree[int, int](nil)
	})

	tree := NewRBTree[int, int](intCompare)
	assert.NotNil(t, tree)
	assert.Equal(t, 0, tree.Len())
}

func TestInsertAndGet(t *testing.T) {
	tree := NewRBTree[int, string](intCompare)
	tree.Insert(5, "five")
	tree.Insert(3, "three")
	tree.Insert(7, "seven")

	val, ok := tree.Get(5)
	assert.True(t, ok)
	assert.Equal(t, "five", val)

	val, ok = tree.Get(3)
	assert.True(t, ok)
	assert.Equal(t, "three", val)

	val, ok = tree.Get(7)
	assert.True(t, ok)
	assert.Equal(t, "seven", val)

	val, ok = tree.Get(10)
	assert.False(t, ok)
	assert.Empty(t, val)

	assert.Equal(t, 3, tree.Len())

	// 更新已有key
	tree.Insert(5, "FIVE")
	val, _ = tree.Get(5)
	assert.Equal(t, "FIVE", val)
}

func TestDelete(t *testing.T) {
	tree := NewRBTree[int, string](intCompare)
	tree.Insert(5, "five")
	tree.Insert(3, "three")
	tree.Insert(7, "seven")

	// 删除不存在的key
	assert.False(t, tree.Delete(10))
	assert.Equal(t, 3, tree.Len())

	// 删除叶子节点
	assert.True(t, tree.Delete(3))
	assert.Nil(t, tree.findNodeByKey(3))
	assert.Equal(t, 2, tree.Len())

	// 删除根节点
	assert.True(t, tree.Delete(5))
	assert.Nil(t, tree.findNodeByKey(5))
	assert.Equal(t, 1, tree.Len())

	// 删除最后一个节点
	assert.True(t, tree.Delete(7))
	assert.Nil(t, tree.findNodeByKey(7))
	assert.Equal(t, 0, tree.Len())
}

func TestLowerAndHigher(t *testing.T) {
	tree := NewRBTree[int, string](intCompare)
	keys := []int{5, 3, 7, 2, 4, 6, 8}
	for _, k := range keys {
		tree.Insert(k, fmt.Sprintf("%d", k))
	}

	val, ok := tree.Lower(5)
	assert.True(t, ok)
	assert.Equal(t, "4", val)

	val, ok = tree.Higher(5)
	assert.True(t, ok)
	assert.Equal(t, "6", val)

	val, ok = tree.Lower(2)
	assert.False(t, ok)

	val, ok = tree.Higher(8)
	assert.False(t, ok)

	val, ok = tree.Lower(10)
	assert.True(t, ok)
	assert.Equal(t, "8", val)

	val, ok = tree.Higher(1)
	assert.True(t, ok)
	assert.Equal(t, "2", val)
}

func TestLen(t *testing.T) {
	tree := NewRBTree[int, string](intCompare)
	assert.Equal(t, 0, tree.Len())

	tree.Insert(5, "five")
	assert.Equal(t, 1, tree.Len())

	tree.Insert(3, "three")
	assert.Equal(t, 2, tree.Len())

	tree.Delete(5)
	assert.Equal(t, 1, tree.Len())

	tree.Delete(3)
	assert.Equal(t, 0, tree.Len())
}

func TestCheck(t *testing.T) {
	tree := NewRBTree[int, string](intCompare)
	tree.Insert(5, "five")
	tree.Insert(3, "three")
	tree.Insert(7, "seven")
	tree.check("After insertions")

	tree.Delete(3)
	tree.check("After deletion of leaf node")

	tree.Delete(7)
	tree.check("After deletion of another node")

	tree.Delete(5)
	tree.check("After deleting root")
}

func TestDeleteOneChildNode(t *testing.T) {
	tree := NewRBTree[int, string](intCompare)

	// 构建一棵树
	keys := []int{5, 3, 7, 6}
	for _, k := range keys {
		tree.Insert(k, fmt.Sprintf("%d", k))
	}

	// 删除有子节点的节点
	assert.True(t, tree.Delete(7))

	// 验证删除后的状态
	val, ok := tree.Get(6)
	assert.True(t, ok)
	assert.Equal(t, "6", val)

	// 验证其他节点仍然存在
	for _, k := range []int{5, 3} {
		val, ok := tree.Get(k)
		assert.True(t, ok)
		assert.Equal(t, fmt.Sprintf("%d", k), val)
	}

	// 验证树的大小
	assert.Equal(t, 3, tree.Len())
}

// 测试 lowestNode 函数
func TestLowestNode(t *testing.T) {
	tree := NewRBTree[int, string](intCompare)
	node := tree.lowestNode()
	assert.Nil(t, node)

	tree.Insert(10, "ten")
	node = tree.lowestNode()
	assert.Equal(t, 10, node.Key)
	assert.Equal(t, "ten", node.Value)

	tree.Insert(5, "five")
	tree.Insert(15, "fifteen")
	tree.Insert(2, "two")
	tree.Insert(7, "seven")
	tree.Insert(20, "twenty")
	node = tree.lowestNode()
	assert.Equal(t, 2, node.Key)
	assert.Equal(t, "two", node.Value)
}

// 测试 highestNode 函数
func TestHighestNode(t *testing.T) {
	tree := NewRBTree[int, string](intCompare)
	node := tree.highestNode()
	assert.Nil(t, node)

	tree.Insert(10, "ten")
	node = tree.highestNode()
	assert.Equal(t, 10, node.Key)
	assert.Equal(t, "ten", node.Value)

	tree.Insert(5, "five")
	tree.Insert(15, "fifteen")
	tree.Insert(2, "two")
	tree.Insert(7, "seven")
	tree.Insert(20, "twenty")
	node = tree.highestNode()
	assert.Equal(t, 20, node.Key)
	assert.Equal(t, "twenty", node.Value)
}

func TestRBTree_Range(t1 *testing.T) {
	tree := NewRBTree[int, int](intCompare)
	for i := 0; i < 10; i++ {
		tree.Insert(i, i)
	}

	var slice []int
	tree.Range(func(key int, value int) bool {
		slice = append(slice, key)
		return true
	})
	assert.Equal(t1, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, slice)
}

func TestRBTree_ReverseRange(t1 *testing.T) {
	tree := NewRBTree[int, int](intCompare)
	for i := 0; i < 10; i++ {
		tree.Insert(i, i)
	}

	var slice []int
	tree.ReverseRange(func(key int, value int) bool {
		slice = append(slice, key)
		return true
	})
	assert.Equal(t1, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, slice)
}
