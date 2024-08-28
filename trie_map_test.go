package triemap

import "testing"

func TestTrieMap(t *testing.T) {
	trieMap := New[string, int]()

	trieMap.Put("apple", 1)
	//trieMap.Put("banana", 2)
	//trieMap.Put("cherry", 3)
	//trieMap.Put("date", 4)
	//trieMap.Put("elderberry", 5)

	apple, ok := trieMap.Get("apple")
	if !ok {
		t.Error("Expected to find apple")
	}

	if apple != 1 {
		t.Error("Expected apple to have value 1, but got", apple)
	}

	t.Log(apple)
}

func TestToBytes(t *testing.T) {
	t.Log(toBytes("apple"))
	a := "apple"
	t.Log(toBytes(a))
}
