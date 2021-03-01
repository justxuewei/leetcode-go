package ds

type Stack struct {
	Arr      []interface{}
	Capacity int
	Len      int
}

func NewStack(capacity int) *Stack {
	stack := &Stack{make([]interface{}, 0, capacity), capacity, 0}
	return stack
}

func (stack *Stack) Pop() interface{} {
	if stack.Len == 0 { return nil }
	lastElem := stack.Arr[stack.Len-1]
	stack.Arr = stack.Arr[:stack.Len-1]
	stack.Len--
	return lastElem
}

func (stack *Stack) Push(element interface{}) {
	if stack.Len >= stack.Capacity { return }
	stack.Arr = append(stack.Arr, element)
	stack.Len++
}
