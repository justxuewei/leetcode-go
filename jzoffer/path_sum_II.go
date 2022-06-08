package jzoffer

func pathSum(root *TreeNode, target int) [][]int {
	var paths [][]int
	var path []int
	dfs3(root, target, 0, &path, &paths)
	return paths
}

func dfs3(node *TreeNode, target, cntsum int, path *[]int, paths *[][]int) {
	if node == nil {
		return
	}
	newsum := cntsum + node.Val
	*path = append(*path, node.Val)
	if isLeaf(node) && newsum == target {
		*paths = append(*paths, deepCopy(*path))
	}
	if !isLeaf(node) {
		dfs3(node.Left, target, newsum, path, paths)
		dfs3(node.Right, target, newsum, path, paths)
	}
	*path = (*path)[:len(*path)-1]
}

func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}

func deepCopy(arr []int) []int {
	newarr := make([]int, 0, len(arr))
	for _, v := range arr {
		newarr = append(newarr, v)
	}
	return newarr
}
