package jzoffer

// https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
// time limit exceeded
func lastRemaining(n int, m int) int {
	m = m % n
	head := &ListNode{Val: 0}
	p := head
	for i := 1; i < n; i++ {
		p.Next = &ListNode{Val: i}
		p = p.Next
	}
	p.Next = head

	for p != p.Next {
		for i := m; i > 1; i-- {
			p = p.Next
		}
		p.Next = p.Next.Next
	}

	return p.Val
}

func lastRemaining1(n int, m int) int {
	ret := 0
	for i := 1; i < n; i++ {
		ret = (ret + m) % (i + 1)
	}
	return ret
}
