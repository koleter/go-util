package tire

// TrieNode 是字典树的节点
type TrieNode[T any] struct {
	children [256]*TrieNode[T]
	isEnd    bool
	val      T
}

// Trie 是字典树
type Trie[T any] struct {
	root *TrieNode[T]
}

// NewTrie 创建一个新的字典树
func NewTrie[T any]() *Trie[T] {
	return &Trie[T]{
		root: &TrieNode[T]{},
	}
}

// Insert 插入一个单词到字典树
func (t *Trie[T]) Insert(word string, val T) {
	t.InsertBytes([]byte(word), val)
}

func (t *Trie[T]) InsertBytes(word []byte, val T) {
	node := t.root

	for _, char := range word {
		if nextNode := node.children[char]; nextNode != nil {
			node = nextNode
		} else {
			newNode := &TrieNode[T]{}
			node.children[char] = newNode
			node = newNode
		}
	}
	node.isEnd = true
	node.val = val
}

// Search 查找一个单词是否在字典树中
func (t *Trie[T]) Search(word string) (T, bool) {
	return t.SearchBytes([]byte(word))
}

func (t *Trie[T]) SearchBytes(word []byte) (T, bool) {
	var zero T
	node := t.root
	for _, char := range word {
		if nextNode := node.children[char]; nextNode != nil {
			node = nextNode
		} else {
			return zero, false
		}
	}
	if node.isEnd {
		return node.val, true
	}
	return zero, false
}

// Match 根据给定的text获取匹配val
func (t *Trie[T]) Match(text string) (T, bool) {
	return t.MatchBytes([]byte(text))
}

func (t *Trie[T]) MatchBytes(word []byte) (T, bool) {
	var zero T
	node := t.root
	for _, char := range word {
		nextNode := node.children[char]
		if nextNode != nil {
			if nextNode.isEnd {
				return nextNode.val, true
			}
			node = nextNode
		} else {
			return zero, false
		}
	}
	return zero, false
}

// MatchLast 根据给定的text获取最长匹配的val
func (t *Trie[T]) MatchLast(text string) (T, bool) {
	return t.MatchLastBytes([]byte(text))
}

func (t *Trie[T]) MatchLastBytes(word []byte) (T, bool) {
	node := t.root
	var res T
	for _, char := range word {
		nextNode := node.children[char]
		if nextNode != nil {
			if nextNode.isEnd {
				res = nextNode.val
			}
			node = nextNode
		} else {
			return res, false
		}
	}
	return res, false
}

// MatchAll 根据给定的text获取所有可匹配的val
func (t *Trie[T]) MatchAll(text string) []T {
	return t.MatchAllBytes([]byte(text))
}

func (t *Trie[T]) MatchAllBytes(word []byte) []T {
	node := t.root
	var res []T
	for _, char := range word {
		nextNode := node.children[char]
		if nextNode != nil {
			if nextNode.isEnd {
				res = append(res, nextNode.val)
			}
			node = nextNode
		} else {
			return res
		}
	}
	return res
}
