package foundation

import (
	"errors"
	ds "github.com/xavier-niu/leetcode/data_struct"
	"log"
)

const treeNodeQueueCapacity = 100

var EmptyTree = errors.New("the amount of tree node should greater than 1")

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// ref: https://stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go
const Null = -int(^uint(0) >> 1) - 1

func GenerateBinaryTree(arr []int) (*TreeNode, error) {
	if len(arr) == 0 { return nil, EmptyTree }

	root := &TreeNode{Val: arr[0]}

	queue := ds.NewQueue(treeNodeQueueCapacity)
	err := queue.Enqueue(root)
	if err != nil {
		return nil, err
	}

	idx := 1
	for queue.Len() > 0 {
		nodeIface, err := queue.Dequeue()
		if err != nil {
			return nil, err
		}

		node, ok := nodeIface.(*TreeNode)
		if !ok {
			log.Fatal("`nodeIface` should be an instance of *TreeNode.")
		}

		// Left
		if len(arr) == idx { break }
		if arr[idx] != Null {
			node.Left = &TreeNode{Val: arr[idx]}
			err = queue.Enqueue(node.Left)
			if err != nil {
				return nil, err
			}
		}
		idx++
		// Right
		if len(arr) == idx { break }
		if arr[idx] != Null {
			node.Right = &TreeNode{Val: arr[idx]}
			err = queue.Enqueue(node.Right)
			if err != nil {
				return nil, err
			}
		}
		idx++
	}

	return root, nil
}
