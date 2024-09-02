package triemap

type TrieMap[K, V any] struct {
	root *trieNode[V]
}

func New[K, V any]() *TrieMap[K, V] {
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
