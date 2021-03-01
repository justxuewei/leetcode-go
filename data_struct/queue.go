package ds

type Queue = []interface{}

func NewQueue(capacity int) Queue {
	queue := make([]interface{}, 0, capacity)
	return queue
}
