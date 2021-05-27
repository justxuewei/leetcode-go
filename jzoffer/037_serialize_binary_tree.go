package jzoffer

import (
	"fmt"
	"strconv"
	"strings"
)

type BTreeCodec struct {}

func NewBTreeCodec() BTreeCodec {
	return BTreeCodec{}
}

func (c *BTreeCodec) serialize(root *TreeNode) string {
	if root == nil { return "" }
	var b strings.Builder
	// queue init
	var queue []*TreeNode
	queue = append(queue, root)
	// last tree node in one level
	l := root
	// level-order traversal
	var node *TreeNode
	for len(queue) > 0 && l != nil {
		node = queue[0]
		queue = queue[1:]
		if node != nil {
			queue = append(queue, node.Left, node.Right)
			_, _ = fmt.Fprintf(&b, "%d,", node.Val)
		} else {
			_, _ = fmt.Fprintf(&b, "n,")
		}
		// replace last tree node
		if node == l {
			for i:=len(queue)-1; i>=0; i-- {
				l = queue[i]
				if queue[i] != nil {
					break
				}
			}
		}
	}
	return b.String()[:b.Len()-1]
}

func (c *BTreeCodec) deserialize(data string) *TreeNode {
	if data == "" { return nil }
	datarr := strings.Split(data, ",")
	var queue []*TreeNode
	root := aToN(datarr[0])
	queue = append(queue, root)
	var node *TreeNode
	for i:=1; i<len(datarr); i+=2 {
		node = queue[0]
		queue = queue[1:]
		if leftNode := aToN(datarr[i]); leftNode != nil {
			node.Left = leftNode
			queue = append(queue, leftNode)
		}
		if i + 1 < len(datarr) {
			if rightNode := aToN(datarr[i+1]); rightNode != nil {
				node.Right = rightNode
				queue = append(queue, rightNode)
			}
		}
	}
	return root
}

func aToN(s string) *TreeNode {
	i, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &TreeNode{Val: i}
}
