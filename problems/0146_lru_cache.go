package problems

// FAILED

type BiLinkNode struct {
	Key, Val int
	Prev, Next *BiLinkNode
}

type LRUCache struct {
	Data map[int]*BiLinkNode
	Head, Tail *BiLinkNode
	Cap int
}


func NewLRUCache(capacity int) LRUCache {
	head := &BiLinkNode{}
	return LRUCache{
		Data: make(map[int]*BiLinkNode, capacity),
		Head: head,
		Tail: head,
		Cap: capacity,
	}
}


func (c *LRUCache) Get(key int) (ret int) {
	if node, ok := c.Data[key]; ok {
		ret = node.Val

		if node.Prev != c.Head {
			if node == c.Tail {
				c.Tail = node.Prev
				node.Prev.Next = nil
			}
			headNext := c.Head.Next
			c.Head.Next = node
			node.Prev = c.Head
			node.Next = headNext
			headNext.Prev = node
		}
		return
	}

	return -1
}


func (c *LRUCache) Put(key int, value int)  {
	if node, ok := c.Data[key]; ok {
		node.Val = value
	}

	headNext := c.Head.Next
	newNode := &BiLinkNode{Key: key, Val: value, Prev: c.Head, Next: headNext}
	c.Head.Next = newNode
	if headNext != nil {
		headNext.Prev = newNode
	}

	if len(c.Data) == c.Cap {
		lastNode := c.Tail
		c.Tail = lastNode.Prev
		c.Tail.Next = nil
		lastNode.Prev = nil
		delete(c.Data, lastNode.Key)
	} else if len(c.Data) == 0 {
		c.Tail = newNode
	}
	
	c.Data[key] = newNode
}
