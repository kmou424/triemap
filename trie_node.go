package triemap

import "sync"

type trieNode[V any] struct {
	children [256]*trieNode[V]

	l sync.RWMutex

	val V
}

func (n *trieNode[V]) getChild(b []byte) (child *trieNode[V]) {
	var walk func(node *trieNode[V], b []byte, idx int)
	walk = func(node *trieNode[V], b []byte, idx int) {
		if idx == len(b) {
			node.l.RLock()
			defer node.l.RUnlock()
			child = node
			return
		}

		c := b[idx]
		if node.children[c] == nil {
			return
		}

		walk(node.children[c], b, idx+1)
	}

	walk(n, b, 0)

	return
}

func (n *trieNode[V]) setChild(b []byte, val V) {
	var walk func(node *trieNode[V], b []byte, idx int)
	walk = func(node *trieNode[V], b []byte, idx int) {
		if idx == len(b) {
			node.l.Lock()
			defer node.l.Unlock()
			node.val = val
			return
		}

		c := b[idx]
		if node.children[c] == nil {
			node.children[c] = &trieNode[V]{}
		}

		walk(node.children[c], b, idx+1)
	}

	walk(n, b, 0)
}
