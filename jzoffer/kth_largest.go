package jzoffer

import "math"

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/

const null = math.MinInt32

func kthLargest(root *TreeNode, k int) int {
	if root == nil { return 0 }
	kth := null
	reversedinorder(root, &k, &kth)
	return kth
}

func reversedinorder(root *TreeNode, k, v *int) {
	if root == nil { return }
	reversedinorder(root.Right, k, v)
	*k--
	if *k == 0 {
		*v = root.Val
		return
	}
	reversedinorder(root.Left, k, v)
}

func kthLargest1(root *TreeNode, k int) int {
	var stack []*TreeNode

	stack = append(stack, root)
	cntnode := root.Right
	for len(stack) > 0 || cntnode != nil {
		if cntnode != nil {
			stack = append(stack, cntnode)
			cntnode = cntnode.Right
			continue
		}
		cntnode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return cntnode.Val
		}
		cntnode = cntnode.Left
	}
	return null
}
