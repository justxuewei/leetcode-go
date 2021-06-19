package ds

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack(10)
	stack.Push(1)
	stack.Push(2)
	for _, v := range stack.arr {
		if intV, ok := v.(int); ok {
			t.Logf("value: %d", intV)
		}
	}
	val, _ := stack.Pop()
	if intV, ok := val.(int); ok {
		t.Logf("popped item is %d", intV)
	}
	for _, v := range stack.arr {
		if intV, ok := v.(int); ok {
			t.Logf("value: %d", intV)
		}
	}
}
