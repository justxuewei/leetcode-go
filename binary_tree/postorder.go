package btree

import (
	"fmt"
	ds "github.com/xavier-niu/leetcode/data_struct"
)

func postorderTraversalRecursively(root *TreeNode, output *[]int) {
	if root == nil { return }
	postorderTraversalRecursively(root.Left, output)
	postorderTraversalRecursively(root.Right, output)
	*output = append(*output, root.Val)
}

type postorderTraversalStackItem struct {
	node *TreeNode
	visitedRight bool
}

func postorderTraversalIteratively(root *TreeNode) []int {
	if root == nil { return nil }

	output := make([]int, 0)
	stack := ds.NewStackWithoutCap()
	crtNode := root
	for crtNode != nil {
		fmt.Println(crtNode)

		topItem, ok := stack.TopElement().(postorderTraversalStackItem)
		if !ok || topItem.node != crtNode {
			stack.Push(postorderTraversalStackItem{node: crtNode, visitedRight: false})
			topItem, _ = stack.TopElement().(postorderTraversalStackItem)
		}
		if !topItem.visitedRight && crtNode.Left != nil {
			crtNode = crtNode.Left
			continue
		}

		if !topItem.visitedRight && crtNode.Right != nil {
			topItem.visitedRight = true
			crtNode = crtNode.Right
			continue
		}
		topItem.visitedRight = true
		output = append(output, crtNode.Val)
		_ = stack.Pop()
		topItem, ok = stack.TopElement().(postorderTraversalStackItem)
		if ok {
			crtNode = topItem.node
		} else {
			crtNode = nil
		}
	}
	return output
}
