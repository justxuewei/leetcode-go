package dp

func generateTrees(n int) []*TreeNode {
	if n == 1 {
		return []*TreeNode{{Val: 1}}
	}

	var dp [][]*TreeNode
	// dp[0], `Val == 0` means the node is empty
	dp = append(dp, []*TreeNode{{Val: 0}})
	// dp[1]
	dp = append(dp, []*TreeNode{{Val: 1}})

	for i:=2; i<=n; i++ {
		dp = append(dp, []*TreeNode{})
		for j:=1; j<=i; j++ {
			var leftTrees, rightTrees []*TreeNode
			leftTrees = dp[j-1]
			for t := range dp[i-j] {
				rightTrees = append(rightTrees, bfsAddOffset(dp[i-j][t], j))
			}
			for _, lt := range leftTrees {
				for _, rt := range rightTrees {
					root := &TreeNode{Val: j}
					if lt.Val != 0 {
						root.Left = lt
					}
					if rt.Val != 0 {
						root.Right = rt
					}
					dp[i] = append(dp[i], root)
				}
			}
		}
	}

	return dp[n]
}

func bfsAddOffset(root *TreeNode, offset int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == 0 {
		return root
	}
	node := &TreeNode{Val: root.Val+offset}
	node.Left = bfsAddOffset(root.Left, offset)
	node.Right = bfsAddOffset(root.Right, offset)
	return node
}
