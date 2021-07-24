package jzoffer1

type CQueue struct {
	in, out []int
}

func NewCQueue() CQueue {
	var in, out []int
	return CQueue{
		in: in,
		out: out,
	}
}

func (q *CQueue) AppendTail(value int)  {
	q.in = append(q.in, value)
}

func (q *CQueue) DeleteHead() (ret int) {
	if len(q.out) == 0 {
		for len(q.in) > 0 {
			q.out = append(q.out, q.in[len(q.in)-1])
			q.in = q.in[:len(q.in)-1]
		}
	}

	if len(q.out) == 0 {
		return -1
	}
	ret = q.out[len(q.out)-1]
	q.out = q.out[:len(q.out)-1]
	return
}
