package linked_list

func hasCycle(head *ListNode) bool {
	if head == nil { return false }
	p1, p2 := head, head
	for p2.Next != nil && p2.Next.Next != nil {
		p1, p2 = p1.Next, p2.Next.Next
		if p1 == p2 {
			return true
		}
	}
	return false
}
