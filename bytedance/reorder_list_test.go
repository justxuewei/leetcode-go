package bytedance

import "testing"

func TestReverseList(t *testing.T) {
	head := &ListNode{Val: 1, Next: nil}
	node1 := &ListNode{Val: 2, Next: nil}
	node2 := &ListNode{Val: 3, Next: nil}
	head.Next = node1
	node1.Next = node2

	head = reverseList(head)
	p := head
	for p != nil {
		t.Log(p.Val)
		p = p.Next
	}
}

func TestReorderList(t *testing.T) {
	head := &ListNode{Val: 1, Next: nil}
	node1 := &ListNode{Val: 2, Next: nil}
	node2 := &ListNode{Val: 3, Next: nil}
	node3 := &ListNode{Val: 4, Next: nil}
	head.Next = node1
	node1.Next = node2
	node2.Next = node3

	reorderList(head)
}
