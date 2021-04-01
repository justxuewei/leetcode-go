package bytedance

import btree "github.com/xavier-niu/leetcode/binary_tree"

func partSumBytedance(node *btree.TreeNode, targetSum int) int {
	npath := 0
	dfs(node, targetSum, 0, &npath)
	return npath
}

func dfs(node *btree.TreeNode, targetSum, currentSum int, num *int) {
	if node == nil { return }
	newSum := currentSum + node.Val
	if node.Left == nil && node.Right == nil {
		if targetSum == newSum {
			*(num)++
		}
		return
	}
	dfs(node.Left, targetSum, newSum, num)
	dfs(node.Right, targetSum, newSum, num)
}
