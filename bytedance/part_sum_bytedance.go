package bytedance

import btree "github.com/xavier-niu/leetcode/btree"

func partSumBytedance(node *btree.TreeNode, targetSum int) int {
	npath := 0
	dfs(node, targetSum, 0, &npath)
	return npath
}

func dfs(node *btree.TreeNode, targetSum, currentSum int, num *int) {
	if node == nil {
		return
	}
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

func dfs1(node *btree.TreeNode, targetSum, currentSum int) int {
	if node == nil {
		return 0
	}
	newSum := currentSum + node.Val
	if node.Left == nil && node.Right == nil && targetSum == newSum {
		return 1
	}
	if targetSum < newSum {
		return 0
	}
	return dfs1(node.Left, targetSum, newSum) + dfs1(node.Right, targetSum, newSum)
}
