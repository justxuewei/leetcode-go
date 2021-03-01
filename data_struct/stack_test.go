package ds

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack(10)
	stack.Push(1)
	stack.Push(2)
	for i, v := range stack.Arr {
		if intV, ok := v.(int); ok {
			t.Logf("index: %d, value: %d", i, intV)
		}
	}
	val := stack.Pop()
	if intV, ok := val.(int); ok {
		t.Logf("popped item is %d", intV)
	}
}
