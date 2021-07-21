package jzoffer1

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	p := root

	var (
		carry bool
		sum int
	)
	for l1 != nil || l2 != nil {
		sum = 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if carry {
			sum++
		}

		if sum < 10 {
			carry = false
		} else {
			sum %= 10
			carry = true
		}

		p.Next = &ListNode{Val: sum}
		p = p.Next
	}

	if carry {
		p.Next = &ListNode{Val: 1}
	}

	return root.Next
}
