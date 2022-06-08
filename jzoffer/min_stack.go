package jzoffer

type MinStack struct {
	RStack []int // regular stack
	DStack []int // descending stack
}

func ConstructorMinStack() MinStack {
	return MinStack{RStack: make([]int, 0), DStack: make([]int, 0)}
}

func (m *MinStack) Push(val int) {
	m.RStack = append(m.RStack, val)

	if len(m.DStack) == 0 {
		m.DStack = append(m.DStack, val)
	} else {
		dtop := m.DStack[len(m.DStack)-1]
		if dtop >= val {
			m.DStack = append(m.DStack, val)
		}
	}
}

func (m *MinStack) Pop() {
	rtop := m.RStack[len(m.RStack)-1]
	m.RStack = m.RStack[:len(m.RStack)-1]
	dtop := m.DStack[len(m.DStack)-1]
	if rtop == dtop {
		m.DStack = m.DStack[:len(m.DStack)-1]
	}
}

func (m *MinStack) Top() int {
	return m.RStack[len(m.RStack)-1]
}

func (m *MinStack) GetMin() int {
	return m.DStack[len(m.DStack)-1]
}
