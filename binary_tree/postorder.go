package btree

import (
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
	stack.Push(&postorderTraversalStackItem{node: root, visitedRight: false})
	crtNode := root.Left

	for {
		topElem, _ := stack.TopElement()
		stackItem, _ := topElem.(*postorderTraversalStackItem)
		if crtNode == nil {
			// set crtNode to the node in the top element of the stack.
			// because stack.Len() > 0, the error could be overlooked.
			crtNode = stackItem.node
		} else {
			// push crtNode to the stack if the node in the top element of the stack isn't equals to crtNode.
			if stackItem.node != crtNode {
				stack.Push(&postorderTraversalStackItem{node: crtNode, visitedRight: false})
				crtNode = crtNode.Left
			} else {
				// if right branch is not visited, visit it right now!
				if !stackItem.visitedRight {
					stackItem.visitedRight = true
					crtNode = crtNode.Right
				} else {
					output = append(output, crtNode.Val)
					_, _ = stack.Pop()
					if stack.Len() == 0 {
						break
					}
					topElem, _ := stack.TopElement()
					stackItem, _ = topElem.(*postorderTraversalStackItem)
					crtNode = stackItem.node
				}
			}
		}
	}
	return output
}
