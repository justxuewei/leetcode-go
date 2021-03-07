package btree

import (
	ds "github.com/xavier-niu/leetcode/data_struct"
)

func preorderTraversalRecursively(root *TreeNode, output []int) []int {
	if root == nil { return output }
	if output == nil {
		output = make([]int, 0)
	}
	output = append(output, root.Val)
	output = preorderTraversalRecursively(root.Left, output)
	output = preorderTraversalRecursively(root.Right, output)
	return output
}

func preorderTraversalIteratively(root *TreeNode) []int {
	if root == nil { return nil }
	output := make([]int, 0)

	stack := ds.NewStackWithoutCap()
	currentNode := root

	for currentNode != nil {
		output = append(output, currentNode.Val)
		stack.Push(currentNode)
		if currentNode.Left == nil {
			// no left node
			var tmpNode *TreeNode
			for {
				tmpNode, _ = stack.Pop().(*TreeNode)
				if tmpNode == nil || tmpNode.Right != nil { break }
			}
			if tmpNode == nil { // there is no more nodes in stack
				break
			}
			currentNode = tmpNode.Right
		} else {
			currentNode = currentNode.Left
		}
	}

	return output
}
