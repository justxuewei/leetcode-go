package jzoffer

type CQueue struct {
	AppendStack []int
	DeleteStack []int
}

func Constructor() CQueue {
	return CQueue{
		AppendStack: make([]int, 0),
		DeleteStack: make([]int, 0),
	}
}

func (queue *CQueue) AppendTail(value int) {
	queue.AppendStack = append(queue.AppendStack, value)
}

func (queue *CQueue) DeleteHead() int {
	if len(queue.AppendStack) == 0 && len(queue.DeleteStack) == 0 {
		return -1
	}
	if len(queue.DeleteStack) == 0 {
		for len(queue.AppendStack) != 0 {
			queue.DeleteStack = append(queue.DeleteStack, queue.AppendStack[len(queue.AppendStack)-1])
			queue.AppendStack = queue.AppendStack[:len(queue.AppendStack)-1]
		}
	}
	ret := queue.DeleteStack[len(queue.DeleteStack)-1]
	queue.DeleteStack = queue.DeleteStack[:len(queue.DeleteStack)-1]
	return ret
}
