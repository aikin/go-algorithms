package trie

import "sync"

type Bytes []byte

type Trie struct {
	rw   sync.RWMutex
	root *trieNode
	size int
}

type trieNode struct {
	children map[byte]*trieNode
	symbol   byte
	value    []byte
	root     bool
}

func NewTrie() *Trie {
	return &Trie{
		root: &trieNode{root: true, children: make(map[byte]*trieNode)},
		size: 1,
	}
}

func newNode(symbol byte) *trieNode {
	return &trieNode{children: make(map[byte]*trieNode), symbol: symbol}
}

func (t *Trie) Size() int {
	t.rw.RLock()
	defer t.rw.RUnlock()
	return t.size
}

func (t *Trie) Insert(key, value Bytes) {
	t.rw.Lock()
	defer t.rw.Unlock()

	currNode := t.root

	for _, symbol := range key {
		if currNode.children[symbol] == nil {
			currNode.children[symbol] = newNode(symbol)
		}

		currNode = currNode.children[symbol]
	}

	// Only increment size if the key value pair is new, otherwise we consider
	// the operation as an update.
	if currNode.value == nil {
		t.size++
	}

	currNode.value = value
}

func (t *Trie) Search(key Bytes) (Bytes, bool) {
	t.rw.RLock()
	defer t.rw.RUnlock()

	currNode := t.root

	for _, symbol := range key {
		if currNode.children[symbol] == nil {
			return nil, false
		}

		currNode = currNode.children[symbol]
	}

	return currNode.value, true
}
