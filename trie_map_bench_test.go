package triemap

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkTrieMap_Put(b *testing.B) {
	trieMap := New[string, int]()
	for i := 0; i < b.N; i++ {
		trieMap.Put(fmt.Sprintf("%d", i), i)
	}
}

func BenchmarkTrieMap_Get(b *testing.B) {
	trieMap := New[string, int]()
	trieMap.Put("key", 1)
	for i := 0; i < b.N; i++ {
		trieMap.Get("key")
	}
}

func BenchmarkSyncMap_Put(b *testing.B) {
	mp := sync.Map{}
	for i := 0; i < b.N; i++ {
		mp.Store(fmt.Sprintf("key%d", i), i)
	}
}

func BenchmarkSyncMap_Get(b *testing.B) {
	mp := sync.Map{}
	mp.Store("key", 1)
	for i := 0; i < b.N; i++ {
		mp.Load("key")
	}
}

func BenchmarkMap_Put(b *testing.B) {
	mp := make(map[string]int)
	for i := 0; i < b.N; i++ {
		mp[fmt.Sprintf("key%d", i)] = i
	}
}

func BenchmarkMap_Get(b *testing.B) {
	mp := make(map[string]int)
	mp["key"] = 1
	for i := 0; i < b.N; i++ {
		_ = mp["key"]
	}
}
