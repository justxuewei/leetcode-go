package btree

import (
	"errors"
	ds "github.com/xavier-niu/leetcode/data_struct"
)

const treeNodeQueueCapacity = 100

var EmptyTree = errors.New("the amount of tree node should greater than 1")

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// ref: https://stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go
const MinInt = -int(^uint(0) >> 1) - 1

func generateTree(arr []int) (*TreeNode, error) {
	if len(arr) == 0 { return nil, EmptyTree }

	root := &TreeNode{Val: arr[0]}

	queue := ds.NewQueue(treeNodeQueueCapacity)
	err := queue.Enqueue(root)
	if err != nil {
		return nil, err
	}

	idx := 1
	for queue.Len() > 0 {
		node, err := queue.Dequeue()
		if err != nil {
			return nil, err
		}

		if len(arr) == idx + 1 {
			break
		}
		if arr[idx] == MinInt {
			
		}

		idx += 2
	}

	return root, nil
}
