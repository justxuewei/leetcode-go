package bytedance

import "math"

type ListNode struct {
	Val int
	Next *ListNode
}

func reorderList(head *ListNode) {
	n, p := 0, head
	for p != nil {
		n++
		p = p.Next
	}

	if n < 3 { return }

	half, counter, p := int(math.Ceil(float64(n)/2.0)), 0, head
	var halfHead *ListNode
	for p != nil {
		if counter == half - 1 {
			halfHead = p.Next
			p.Next = nil
			break
		}
		p = p.Next
		counter++
	}
	halfHead = reverseList(halfHead)

	p, hp := head, halfHead
	for p != nil {
		pNext := p.Next
		p.Next = hp
		p = pNext
		if hp != nil {
			hpNext := hp.Next
			hp.Next = pNext
			hp = hpNext
		}
	}
}

func reverseList(head *ListNode) *ListNode {
	p, n := head, head.Next
	p.Next = nil
	for n != nil {
		n2 := n.Next
		n.Next = p
		n, p = n2, n
	}
	return p
}
