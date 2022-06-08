package jzoffer

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	p1, p2 := head, head.Next
	// reverse the linked list
	head.Next = nil
	length := 0
	for p2 != nil {
		p3 := p2.Next
		p2.Next = p1
		p1, p2 = p2, p3
		length++
	}
	ret := make([]int, 0, length)
	for p1 != nil {
		ret = append(ret, p1.Val)
		p1 = p1.Next
	}
	return ret
}
