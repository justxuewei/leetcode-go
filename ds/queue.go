package ds

// var QueueNoSpace = errors.New("the length of the queue exceeds the capacity")
// var QueueNoElem = errors.New("there is no elements")

type Queue struct {
	arr  []interface{}
	head int
	tail int
}

func NewQueue(length int) *Queue {
	queue := &Queue{
		make([]interface{}, length+1),
		0,
		0,
	}
	return queue
}

func (queue *Queue) Enqueue(elem interface{}) error {
	nextTail := (queue.tail + 1) % len(queue.arr)
	if nextTail == queue.head {
		return NoSpace
	}
	queue.arr[queue.tail] = elem
	queue.tail = nextTail
	return nil
}

func (queue *Queue) Dequeue() (interface{}, error) {
	if queue.head == queue.tail {
		return nil, NoElement
	}
	elem := queue.arr[queue.head]
	queue.head = (queue.head + 1) % len(queue.arr)
	return elem, nil
}

func (queue *Queue) Len() int {
	head := queue.head
	tail := queue.tail
	if head > tail {
		tail += len(queue.arr)
	}
	return tail - head
}

func (queue *Queue) Tail() interface{} {
	return queue.arr[(queue.tail-1+len(queue.arr))%len(queue.arr)]
}
