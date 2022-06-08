package jzoffer

// ---- Queue ----

type Queue struct {
	head, tail *queueNode
}

func NewQueue() *Queue {
	head := &queueNode{}
	tail := head
	return &Queue{
		head: head,
		tail: tail,
	}
}

func (q *Queue) Enqueue(node *TreeNode) {
	if node == nil {
		return
	}
	q.tail.next = &queueNode{Node: node, prev: q.tail}
	q.tail = q.tail.next
}

func (q *Queue) Dequeue() (node *TreeNode) {
	if q.head.next == nil {
		return nil
	}
	node = q.head.next.Node
	q.head.next = q.head.next.next
	if q.head.next == nil {
		q.tail = q.head
	}
	return
}

func (q *Queue) Head() *TreeNode {
	if q.head.next == nil {
		return nil
	}
	return q.head.next.Node
}

// ---- QueueNode ----

type queueNode struct {
	Node       *TreeNode
	prev, next *queueNode
}

// ---- Main ----

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	queue := NewQueue()
	var ret []int

	queue.Enqueue(root)

	for queue.Head() != nil {
		treeNode := queue.Dequeue()
		ret = append(ret, treeNode.Val)
		queue.Enqueue(treeNode.Left)
		queue.Enqueue(treeNode.Right)
	}

	return ret
}
