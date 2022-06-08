package jzoffer

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return nil
	}
	p1, p2 := head, head
	for i := 0; i < k; i++ {
		if p2 == nil {
			return nil
		}
		p2 = p2.Next
	}

	for p2 != nil {
		p1, p2 = p1.Next, p2.Next
	}

	return p1
}
