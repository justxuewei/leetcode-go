package jzoffer1

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	p := head

	var newhead *Node

	var newnode *Node
	for p != nil {
		newnode = &Node{Val: p.Val, Next: p.Next, Random: p.Random}
		p.Next = newnode
		p = p.Next.Next
	}

	newhead = head.Next

	p = head
	for p != nil {
		if p.Next.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}

	p = head
	for p != nil {
		newnode = p.Next
		p.Next = newnode.Next
		p = newnode.Next
		if newnode.Next != nil {
			newnode.Next = newnode.Next.Next
		}
	}

	return newhead
}
