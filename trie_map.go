package triemap

import (
	"bytes"
	"encoding/gob"
)

type TrieMap[K, V any] struct {
	root *trieNode[V]
}

func New[K, V any]() *TrieMap[K, V] {
	{
		var key K
		gob.Register(key)
	}
	return &TrieMap[K, V]{
		root: &trieNode[V]{},
	}
}

func (t *TrieMap[K, V]) Get(key K) (val V, ok bool) {
	kBytes := toBytes(key)
	if kBytes == nil {
		ok = false
		return
	}

	node := t.root.getChild(kBytes)
	if node == nil {
		ok = false
		return
	}
	return node.val, true
}

func (t *TrieMap[K, V]) Put(key K, val V) {
	kBytes := toBytes(key)
	if kBytes == nil {
		return
	}

	t.root.setChild(kBytes, val)
}

func (t *TrieMap[K, V]) toBytes(key K) []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(key)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}
