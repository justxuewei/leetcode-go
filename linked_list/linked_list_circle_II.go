package linked_list

func detectCycle(head *ListNode) *ListNode {
	if isCircle, p1 := hasCircle1(head); isCircle {
		p2 := head
		for p1 != p2 {
			p1, p2 = p1.Next, p2.Next
		}
		return p1
	}
	return nil
}

func hasCircle1(head *ListNode) (bool, *ListNode) {
	if head == nil || head.Next == nil { return false, nil }
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true, slow
		}
	}
	return false, nil
}
