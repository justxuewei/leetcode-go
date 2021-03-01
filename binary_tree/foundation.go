package btree

import ds "github.com/xavier-niu/leetcode/data_struct"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// ref: https://stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go
const MinInt = -int(^uint(0) >> 1) - 1

func generateTree(arr []int) *TreeNode {
	if len(arr) == 0 { return nil }

	root := &TreeNode{Val: arr[0]}
	if len(arr) == 1 { return root }

	stack := ds.NewStack(100)
	stack.Push(root)

	idx := 1
	for idx < len(arr) {

	}

	return root
}
