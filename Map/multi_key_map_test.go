package Map

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// 测试newNode函数
func TestNewNode(t *testing.T) {
	n := newNode[int, string]()
	if n == nil {
		t.Error("newNode() 返回了nil")
	}
	if n.children == nil {
		t.Error("newNode() 没有正确初始化children map")
	}
	if n.hasVal {
		t.Error("newNode() 初始化的节点hasVal应为false")
	}
}

// 测试getAllValues方法
func TestNode_getAllValues(t *testing.T) {
	// 创建测试结构
	/*
		root
		|
		1
		|
		2 -> val="b"
		|
		3 -> val="a"
	*/
	root := newNode[int, string]()
	node1 := newNode[int, string]()
	node2 := newNode[int, string]()
	node3 := newNode[int, string]()

	root.children[1] = node1
	node1.children[2] = node2
	node2.children[3] = node3

	node2.hasVal = true
	node2.val = "b"

	node3.hasVal = true
	node3.val = "a"

	// 测试获取node2的所有值
	var res []string
	node2.getAllValues(&res)
	if len(res) != 2 || res[0] != "b" || res[1] != "a" {
		t.Errorf("getAllValues() 返回了错误的值: %v", res)
	}

	// 测试获取叶子节点的值
	res = []string{}
	node3.getAllValues(&res)
	if len(res) != 1 || res[0] != "a" {
		t.Errorf("getAllValues() 返回了错误的值: %v", res)
	}

	// 测试没有值的节点
	res = []string{}
	root.getAllValues(&res)
	if len(res) != 2 || res[0] != "b" || res[1] != "a" {
		t.Errorf("getAllValues() 返回了错误的值: %v", res)
	}
}

// 测试NewMultiKeyMap函数
func TestNewMultiKeyMap(t *testing.T) {
	m := NewMultiKeyMap[int, string]()
	if m == nil {
		t.Error("NewMultiKeyMap() 返回了nil")
	}
	if m.root == nil {
		t.Error("NewMultiKeyMap() 没有正确初始化root节点")
	}
}

// 测试Insert和Get方法
func TestMultiKeyMap_InsertAndGet(t *testing.T) {
	m := NewMultiKeyMap[string, int]()

	// 测试基本插入和获取
	keys1 := []string{"a", "b", "c"}
	val1 := 1
	m.Put(keys1, val1)

	res, ok := m.Get(keys1)
	if !ok || res != val1 {
		t.Errorf("Get(%v) 返回了错误的值: %v, %v", keys1, res, ok)
	}

	// 测试覆盖值
	val2 := 2
	m.Put(keys1, val2)

	res, ok = m.Get(keys1)
	if !ok || res != val2 {
		t.Errorf("Get(%v) 返回了错误的值: %v, %v", keys1, res, ok)
	}

	// 测试不存在的键
	keys2 := []string{"a", "b", "d"}
	res, ok = m.Get(keys2)
	if ok {
		t.Errorf("Get(%v) 应该返回false，但返回了true", keys2)
	}

	// 测试空键
	keys3 := []string{}
	m.Put(keys3, val1)
	res, ok = m.Get(keys3)
	if !ok || res != val1 {
		t.Errorf("Get(%v) 返回了错误的值: %v, %v", keys3, res, ok)
	}
}

// 测试GetPrefix方法
func TestMultiKeyMap_GetPrefix(t *testing.T) {
	m := NewMultiKeyMap[int, string]()

	// 插入多个值构建如下结构:
	/*
		root
		|
		1 -> "a"
		|
		2 -> "b"
		|    \
		|     4 -> "d"
		|
		3 -> "c"
	*/
	m.Put([]int{1}, "a")
	m.Put([]int{1, 2}, "b")
	m.Put([]int{1, 2, 4}, "d")
	m.Put([]int{1, 3}, "c")

	// 测试获取前缀[1]
	res := m.GetPrefix([]int{1})
	expected := []string{"a", "b", "d", "c"}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("GetPrefix([1]) 返回了错误的结果: %v, 期望: %v", res, expected)
	}

	// 测试获取前缀[1,2]
	res = m.GetPrefix([]int{1, 2})
	expected = []string{"b", "d"}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("GetPrefix([1,2]) 返回了错误的结果: %v, 期望: %v", res, expected)
	}

	// 测试不存在的前缀
	res = m.GetPrefix([]int{1, 5})
	expected = nil
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("GetPrefix([1,5]) 返回了错误的结果: %v, 期望: %v", res, expected)
	}

	// 测试空前缀
	res = m.GetPrefix([]int{})
	expected = []string{"a", "b", "d", "c"}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("GetPrefix([]) 返回了错误的结果: %v, 期望: %v", res, expected)
	}
}

func TestMultiKeyMap_GetPrefix2(t *testing.T) {
	m := NewMultiKeyMap[string, string]()
	m.Put([]string{"mongo", "1.0.0.1"}, "1")
	m.Put([]string{"mongo", "1.0.0.2"}, "2")
	strings := m.GetPrefix([]string{"mongo"})
	assert.Equal(t, strings[0], "1")
	assert.Equal(t, strings[1], "2")
}

// 测试多层路径插入和获取
func TestMultiKeyMap_MultiLevelPath(t *testing.T) {
	m := NewMultiKeyMap[int, string]()

	// 插入多层路径
	keys := []int{1, 2, 3, 4, 5}
	val := "test"
	m.Put(keys, val)

	// 获取完整路径
	res, ok := m.Get(keys)
	if !ok || res != val {
		t.Errorf("Get(%v) 返回了错误的值: %v, %v", keys, res, ok)
	}

	// 获取中间路径
	shorterKeys := []int{1, 2, 3}
	res, ok = m.Get(shorterKeys)
	if ok {
		t.Errorf("Get(%v) 应该没有值，但返回了: %v, %v", shorterKeys, res, ok)
	}

	// 获取中间路径的前缀
	prefixRes := m.GetPrefix(shorterKeys)
	assert.NotEqual(t, 0, len(prefixRes))
}

// 测试并发操作
func TestMultiKeyMap_ConcurrentAccess(t *testing.T) {
	m := NewMultiKeyMap[string, int]()

	// 并发插入不同键
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(i int) {
			keys := []string{"prefix", fmt.Sprintf("%d", i)}
			m.Put(keys, i)
			done <- true
		}(i)
	}

	// 等待所有goroutine完成
	for i := 0; i < 10; i++ {
		<-done
	}

	// 验证所有值
	for i := 0; i < 10; i++ {
		keys := []string{"prefix", fmt.Sprintf("%d", i)}
		res, ok := m.Get(keys)
		if !ok || res != i {
			t.Errorf("Get(%v) 返回了错误的值: %v, %v", keys, res, ok)
		}
	}
}

func TestMultiKeyMap_Delete(t *testing.T) {
	m := NewMultiKeyMap[int, string]()
	keys := []int{1, 2, 3, 4, 5}
	value := "value"
	m.Put(keys, value)
	get, exist := m.Get(keys)
	assert.True(t, exist)
	assert.Equal(t, value, get)
	m.Delete(keys)
	get, exist = m.Get(keys)
	assert.False(t, exist)
}
