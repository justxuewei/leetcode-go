package ds

type Stack struct {
	arr []interface{}
}

func NewStackWithoutCap() *Stack {
	return NewStack(0)
}

func NewStack(capacity int) *Stack {
	stack := &Stack{make([]interface{}, 0, capacity)}
	return stack
}

func (stack *Stack) Pop() interface{} {
	if stack.Len() == 0 { return nil }
	lastElem := stack.arr[stack.Len()-1]
	stack.arr = stack.arr[:stack.Len()-1]
	return lastElem
}

func (stack *Stack) Push(element interface{}) {
	stack.arr = append(stack.arr, element)
}

func (stack *Stack) TopElement() interface{} {
	if stack.Len() == 0 { return nil }
	return stack.arr[stack.Len()-1]
}

func (stack *Stack) Len() int {
	return len(stack.arr)
}
