package jzoffer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	nodemap := make(map[*Node]*Node, 0)
	var newhead, newp *Node
	p := head
	for p != nil {
		newnode := &Node{Val: p.Val}
		if newhead == nil {
			newhead = newnode
		} else {
			newp.Next = newnode
		}
		newp = newnode
		nodemap[p] = newnode
		p = p.Next
	}

	p, newp = head, newhead
	for p != nil {
		if newnoderam, ok := nodemap[p.Random]; ok {
			newp.Random = newnoderam
		}
		p, newp = p.Next, newp.Next
	}

	return newhead
}

func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}

	p := head
	for p != nil {
		newnode := &Node{Val: p.Val, Next: p.Next}
		p.Next = newnode
		p = newnode.Next
	}

	p = head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}

	newhead, newp := head.Next, head.Next
	p = head
	for newp.Next != nil {
		p.Next, newp.Next = p.Next.Next, newp.Next.Next
		p, newp = p.Next, newp.Next
	}
	p.Next = nil

	return newhead
}
