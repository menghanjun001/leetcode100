package leetcode

import "testing"

func TestLRUCache(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(1,1)
	lruCache.Put(1,2)
	lruCache.Get(1)
}
