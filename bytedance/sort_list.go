package bytedance

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { return head }
	middleNode := middleNode(head)
	right := sortList(middleNode.Next)
	middleNode.Next = nil
	left := sortList(head)
	return mergeLists(left, right)
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1, p2 := head, head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return p1
}

func mergeLists(a *ListNode, b *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	p1, p2, p := a, b, head
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}

	if p2 != nil {
		p.Next = p2
	}

	return head.Next
}

//func sortList(head *ListNode) *ListNode {
//	if head == nil { return nil }
//	nsort := 2
//	n := 0
//	p, retHead := head, head
//
//	for p != nil {
//		n++
//		p = p.Next
//	}
//
//	for nsort < n {
//		nsection := nsort/2
//		p = head
//		includeHead := false
//		var p1, p2, ph1, ph2, lastSectionP *ListNode
//		counter := 0
//		for p != nil {
//			nextP := p.Next
//			if ph1 == nil && ph2 == nil && counter == 0 {
//				p1, ph1 = p, p
//				if p == retHead {
//					includeHead = true
//				}
//				counter++
//			} else if ph1 != nil && ph2 == nil && counter > 0 && counter < nsection {
//				p1.Next, p1 = p, p
//				counter++
//			} else if ph1 != nil && ph2 == nil && counter > 0  && counter == nsection {
//				p1.Next = nil
//				nextP = p
//				counter = 0
//			} else if ph1 != nil && ph2 == nil && counter == 0 {
//				p2, ph2 = p, p
//				counter++
//			} else if ph1 != nil && ph2 != nil && counter > 0  && counter < nsection {
//				p2.Next, p2 = p, p
//				counter++
//			} else if ph1 != nil && ph2 != nil && counter > 0 && counter == nsection {
//				p2.Next = nil
//				nextP = p
//				h, l := combineList(ph1, ph2)
//				if includeHead {
//					retHead = h
//					includeHead = false
//				} else {
//					lastSectionP.Next = h
//				}
//				l.Next = nextP
//				counter = 0
//				lastSectionP = l
//				ph1, p1, ph2, p2 = nil, nil, nil, nil
//			}
//			p = nextP
//		}
//		nsort *= 2
//	}
//
//	return retHead
//}
//
//func combineList(h1 *ListNode, h2 *ListNode) (head *ListNode, last *ListNode) {
//	var pret *ListNode
//	p1, p2 := h1, h2
//
//	if h1 == nil {
//		head = h2
//		for p2.Next != nil {
//			p2 = p2.Next
//		}
//		last = p2
//		return head, last
//	}
//
//	if h2 == nil {
//		head = p1
//		for p1.Next != nil {
//			p1 = p1.Next
//		}
//		last = p1
//		return head ,last
//	}
//
//	for p1 != nil && p2 != nil {
//		if p1.Val < p2.Val {
//			if pret == nil {
//				head, pret = p1, p1
//			} else {
//				pret.Next, pret = p1, p1
//			}
//			p1 = p1.Next
//		} else {
//			if pret == nil {
//				head, pret = p2, p2
//			} else {
//				pret.Next, pret = p2, p2
//			}
//			p2 = p2.Next
//		}
//	}
//	if p1 != nil && pret != nil {
//		pret.Next = p1
//	}
//
//	if p2 != nil && pret != nil {
//		pret.Next = p2
//	}
//
//	for pret != nil && pret.Next != nil {
//		pret = pret.Next
//	}
//
//	last = pret
//
//	return head, last
//}
