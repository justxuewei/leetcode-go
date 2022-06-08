package jzoffer

// https://leetcode.com/problems/reverse-linked-list/
// https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p1, p2 := head, head.Next
	p1.Next = nil
	for p2 != nil {
		p3 := p2.Next
		p2.Next = p1
		p1, p2 = p2, p3
	}
	return p1
}
