package bytedance

type LRUCacheNode struct {
	Key  int
	Val  int
	next *LRUCacheNode
	prev *LRUCacheNode
}

type LRUCache struct {
	Head       *LRUCacheNode
	Tail       *LRUCacheNode
	AddressMap map[int]*LRUCacheNode
	Length     int
	Capacity   int
}

func Constructor(capacity int) LRUCache {
	if capacity == 0 {
		panic("The capacity should be greater than 0.")
	}
	amap := make(map[int]*LRUCacheNode, capacity)
	return LRUCache{
		Head:       nil,
		Tail:       nil,
		AddressMap: amap,
		Length:     0,
		Capacity:   capacity,
	}
}

func (cache *LRUCache) Get(key int) int {
	if nodeArr, ok := cache.AddressMap[key]; ok {
		cache.moveToHead(nodeArr)
		return nodeArr.Val
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	if cache.Length > 0 {
		// cache is not empty
		if curNode, ok := cache.AddressMap[key]; ok {
			// cache hit
			if curNode.Val != value {
				curNode.Val = value
			}
			cache.moveToHead(curNode)
			return
		} else {
			// cache missing
			if cache.Length == cache.Capacity {
				// reached capacity of cache
				tailCandidate := cache.Tail.prev
				delete(cache.AddressMap, cache.Tail.Key)
				cache.Tail = nil
				cache.Tail = tailCandidate
				cache.Length--
			}
		}
	}

	// create a node at head
	node := LRUCacheNode{Key: key, Val: value, prev: nil, next: cache.Head}
	if cache.Head != nil {
		cache.Head.prev = &node
	}
	cache.Head = &node
	// update map
	cache.AddressMap[key] = &node
	if cache.Length == 0 {
		cache.Tail = &node
	}
	cache.Length++
}

func (cache *LRUCache) moveToHead(node *LRUCacheNode) {
	if node == cache.Head {
		return
	}
	if node == cache.Tail {
		cache.Tail = cache.Tail.prev
		cache.Tail.next = nil
	}
	prevNode := (*node).prev
	nextNode := (*node).next
	if prevNode != nil {
		prevNode.next = nextNode
	}
	if nextNode != nil {
		nextNode.prev = prevNode
	}
	cache.Head.prev = node
	node.next = cache.Head
	node.prev = nil
	cache.Head = node
}
