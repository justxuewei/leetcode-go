package jzoffer

type MaxQueue struct {
	data, max []int
}


func MaxQueueConstructor() MaxQueue {
	return MaxQueue{
		data: make([]int, 0),
		max: make([]int, 0),
	}
}


func (q *MaxQueue) MaxValue() int {
	if len(q.max) == 0 {
		return -1
	}
	return q.max[0]
}


func (q *MaxQueue) PushBack(value int)  {
	q.data = append(q.data, value)
	if len(q.max) != 0 {
		i := len(q.max)-1
		for ; i>=0 && q.max[i] < value; i-- {}
		q.max = q.max[:i+1]
	}
	q.max = append(q.max, value)
}


func (q *MaxQueue) PopFront() (val int) {
	if len(q.data) == 0 || len(q.max) == 0 {
		return -1
	}
	val = q.data[0]
	q.data = q.data[1:]
	if val == q.max[0] {
		q.max = q.max[1:]
	}
	return
}
