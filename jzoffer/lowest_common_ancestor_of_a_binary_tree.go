package jzoffer

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var pathp, pathq []*TreeNode
	dfs(root, p, &pathp)
	dfs(root, q, &pathq)
	n := min(len(pathp), len(pathq))
	for i := 1; i < n; i++ {
		if pathp[i] != pathq[i] {
			return pathp[i-1]
		}
	}
	return pathp[n-1]
}

func dfs(root, node *TreeNode, path *[]*TreeNode) bool {
	if root == nil {
		return false
	}

	*path = append(*path, root)

	if node == root {
		return true
	}

	if dfs(root.Left, node, path) {
		return true
	}
	if dfs(root.Right, node, path) {
		return true
	}
	// remove history
	*path = (*path)[:len(*path)-1]
	return false
}

func min(lhs, rhs int) int {
	if lhs < rhs {
		return lhs
	}
	return rhs
}

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else {
		return right
	}
}
