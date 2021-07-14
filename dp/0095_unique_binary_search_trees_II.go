package dp

import "github.com/xavier-niu/leetcode/ds"

func generateTrees(n int) []*ds.TreeNode {
	if n == 1 {
		return []*ds.TreeNode{{Val: 1}}
	}

	var dp [][]*ds.TreeNode
	// dp[0]
	dp = append(dp, []*ds.TreeNode{})
	// dp[1]
	dp = append(dp, []*ds.TreeNode{{Val: 1}})

	for i:=2; i<=n; i++ {
		for j:=1; j<=i; j++ {
			// root == j
			// ?
			//nleft, nright := j-1, i-j
			//for l:=0; l<nleft; l++ {
			//	for r:=nright;
			//}
		}
	}

	return dp[n]
}
