package jzoffer

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil { return nil }
	prep, p := head, head
	for p.Val != val && p != nil {
		prep = p
		p = p.Next
	}

	if p == head {
		return head.Next
	}

	if p == nil || p.Next == nil {
		prep.Next = nil
	} else {
		prep.Next = p.Next
	}

	return head
}
