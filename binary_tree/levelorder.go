package btree

import (
	ds "github.com/xavier-niu/leetcode/data_struct"
)

// My Solution: BFS
func levelOrder(root *TreeNode) [][]int {
	if root == nil { return nil}

	var output [][]int

	queue := ds.NewQueue(50)
	_ = queue.Enqueue(root)

	// current level
	level := 0
	output = append(output, make([]int, 0))
	// a tree node indicated the corresponding level is end.
	levelEndNode := root

	for {
		head, _ := queue.Dequeue()
		headElem := head.(*TreeNode)
		output[level] = append(output[level], headElem.Val)
		if headElem.Left != nil {
			_ = queue.Enqueue(headElem.Left)
		}
		if headElem.Right != nil {
			_ = queue.Enqueue(headElem.Right)
		}

		if queue.Len() == 0 { break }
		if levelEndNode == headElem {
			level++
			levelEndNode = queue.Tail().(*TreeNode)
			output = append(output, make([]int, 0))
		}
	}

	return output
}
