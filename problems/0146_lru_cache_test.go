package problems

import "testing"

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(4, 1)
	t.Log(cache.Get(1))
	t.Log(cache.Get(2))
}
