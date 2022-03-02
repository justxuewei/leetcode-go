package jzoffer2

func detectCycle(head *ListNode) *ListNode {
	pslow, pfast := head, head
	for pfast != nil && pfast.Next != nil {
		pslow = pslow.Next
		pfast = pfast.Next.Next

		if pslow == pfast {
			p := head
			for p != pslow {
				p = p.Next
				pslow = pslow.Next
			}
			return p
		}
	}
	return nil
}
