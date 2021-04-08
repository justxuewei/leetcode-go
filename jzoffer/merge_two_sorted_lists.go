package jzoffer

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	p1, p2 := l1, l2
	var head, p *ListNode

	for p1 != nil && p2 != nil {
		var next *ListNode

		if p1.Val < p2.Val {
			next = p1
			p1 = p1.Next
		} else {
			next = p2
			p2 = p2.Next
		}

		if head == nil {
			head, p = next, next
		} else {
			p.Next, p = next, next
		}
	}

	if p1 != nil {
		p.Next = p1
	} else {
		p.Next = p2
	}

	return head
}
