package jzoffer1

import "fmt"

func pathSum(root *TreeNode, target int) [][]int {
	var (
		ret [][]int
		path []int
	)
	pathSumBacktrace(root, 0, target, path, &ret)
	fmt.Println(ret)
	return ret
}

func pathSumBacktrace(root *TreeNode, sum, target int, path []int, ret *[][]int) {
	if root == nil {
		return
	}

	path = append(path, root.Val)
	sum += root.Val
	if sum == target && root.Left == nil && root.Right == nil {
		*ret = append(*ret, deepCopy(path))
		return
	}

	pathSumBacktrace(root.Left, sum, target, path, ret)
	pathSumBacktrace(root.Right, sum, target, path, ret)
}

func deepCopy(arr []int) []int {
	ret := make([]int, len(arr))
	for i := range arr {
		ret[i] = arr[i]
	}
	return ret
}
