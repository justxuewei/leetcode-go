package jzoffer

import "math"

func isBalanced(root *TreeNode) bool {
	return dfs2(root, 1) != -1
}

func dfs2(node *TreeNode, height int) int {
	if node == nil { return height }
	lefth := dfs2(node.Left, height+1)
	righth := dfs2(node.Right, height+1)

	if lefth == -1 || righth == -1 {
		return -1
	}

	if int(math.Abs(float64(lefth-righth))) > 1 {
		return -1
	}

	return int(math.Max(float64(lefth), float64(righth)))
}
