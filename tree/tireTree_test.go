package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	trie := NewTrie[int]()
	trie.Insert("44", 44)
	trie.Insert("123", 123)
	search, b := trie.Search("44")
	assert.True(t, b)
	assert.Equal(t, 44, search)
	search, b = trie.Search("123")
	assert.True(t, b)
	assert.Equal(t, 123, search)
	search, b = trie.Search("987")
	assert.False(t, b)
}

func TestTrie_InsertBytes(t *testing.T) {
	trie := NewTrie[int]()
	trie.InsertBytes([]byte("44"), 44)
	trie.InsertBytes([]byte("123"), 123)
	search, b := trie.Search("44")
	assert.True(t, b)
	assert.Equal(t, 44, search)
	search, b = trie.Search("123")
	assert.True(t, b)
	assert.Equal(t, 123, search)
	search, b = trie.Search("987")
	assert.False(t, b)
}

func TestTrie_Match(t *testing.T) {
	trie := NewTrie[int]()
	trie.Insert("hello", 1)
	_, b := trie.Match("hello world")
	assert.True(t, b)
	_, b = trie.Match("hello lisi")
	assert.True(t, b)
	_, b = trie.Match("aaa")
	assert.False(t, b)
}

func TestTrie_MatchBytes(t *testing.T) {
	trie := NewTrie[int]()
	trie.Insert("hello", 1)
	_, b := trie.MatchBytes([]byte("hello world"))
	assert.True(t, b)
	_, b = trie.MatchBytes([]byte("hello lisi"))
	assert.True(t, b)
	_, b = trie.MatchBytes([]byte("aaa"))
	assert.False(t, b)
}

func TestTrie_MatchLast(t *testing.T) {
	trie := NewTrie[int]()
	trie.Insert("hello", 1)
	trie.Insert("hello world", 2)
	last, exist := trie.MatchLast("hello zhangsan")
	assert.Equal(t, 1, last)
	assert.True(t, exist)
	last, exist = trie.MatchLast("hello world!!!!!")
	assert.Equal(t, 2, last)
	assert.True(t, exist)
	_, exist = trie.MatchLast("aaa")
	assert.False(t, exist)
}

func TestTrie_MatchLastBytes(t *testing.T) {
	trie := NewTrie[int]()
	trie.Insert("hello", 1)
	trie.Insert("hello world", 2)
	last, exist := trie.MatchLastBytes([]byte("hello"))
	assert.Equal(t, 1, last)
	assert.True(t, exist)
	last, exist = trie.MatchLastBytes([]byte("hello world!!!!!"))
	assert.Equal(t, 2, last)
	assert.True(t, exist)
	_, exist = trie.MatchLastBytes([]byte("aaa"))
	assert.False(t, exist)
}

func TestTrie_MatchAll(t *testing.T) {
	trie := NewTrie[int]()
	trie.Insert("dog", 2)
	trie.Insert("doghouse", 20)
	all := trie.MatchAll("doghouse so big")
	assert.Equal(t, []int{2, 20}, all)
}

func TestTrie_MatchAllBytes(t *testing.T) {
	trie := NewTrie[int]()
	trie.Insert("dog", 2)
	trie.Insert("doghouse", 20)
	all := trie.MatchAllBytes([]byte("doghouse so big"))
	assert.Equal(t, []int{2, 20}, all)
}

func TestTrie_InsertMultiTimes(t *testing.T) {
	trie := NewTrie[int]()
	trie.Insert("dog", 2)
	search, _ := trie.Search("dog")
	assert.Equal(t, 2, search)
	trie.Insert("dog", 20)
	search, _ = trie.Search("dog")
	assert.Equal(t, 20, search)
}
