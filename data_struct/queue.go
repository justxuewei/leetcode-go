package ds

import "errors"

var QueueNoSpace = errors.New("the length of the queue exceeds the capacity")
var QueueNoElem = errors.New("there is no elements")

type Queue struct {
	arr []interface{}
	head int
	tail int
}

func NewQueue(length int) *Queue {
	queue := &Queue{
		make([]interface{}, length),
		0,
		0,
	}
	return queue
}

func (queue *Queue) Enqueue(elem interface{}) error {
	nextTail := (queue.tail + 1) % len(queue.arr)
	if nextTail == queue.head {
		return QueueNoSpace
	}
	queue.arr[queue.tail] = elem
	queue.tail = nextTail
	return nil
}

func (queue *Queue) Dequeue() (interface{}, error) {
	if queue.head == queue.tail {
		return nil, QueueNoElem
	}
	elem := queue.arr[queue.head]
	queue.head = (queue.head + 1) % len(queue.arr)
	return elem, nil
}
