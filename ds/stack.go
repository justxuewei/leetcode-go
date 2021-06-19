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

func (stack *Stack) Pop() (interface{}, error) {
	if stack.Len() == 0 { return nil, NoElement }
	lastElem := stack.arr[stack.Len()-1]
	stack.arr = stack.arr[:stack.Len()-1]
	return lastElem, nil
}

func (stack *Stack) Push(element interface{}) {
	stack.arr = append(stack.arr, element)
}

func (stack *Stack) TopElement() (interface{}, error) {
	if stack.Len() == 0 { return nil, NoElement }
	return stack.arr[stack.Len()-1], nil
}

func (stack *Stack) Len() int {
	return len(stack.arr)
}
