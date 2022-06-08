package lnklist

type MyLinkedList struct {
	head   *ListNode
	length int
}

func Constructor() MyLinkedList {
	return MyLinkedList{head: nil, length: 0}
}

func (list *MyLinkedList) Get(index int) int {
	if index < 0 || index >= list.length {
		return -1
	}
	counter := 0
	p := list.head
	for counter != index {
		p = p.Next
		counter++
	}
	return p.Val
}

func (list *MyLinkedList) AddAtHead(val int) {
	var node *ListNode
	if list.length == 0 {
		node = &ListNode{Val: val}
	} else {
		node = &ListNode{Val: val, Next: list.head}
	}
	list.head = node
	list.length++
}

func (list *MyLinkedList) AddAtTail(val int) {
	node := &ListNode{Val: val}
	if list.length == 0 {
		list.head = node
	} else {
		p := list.head
		for p.Next != nil {
			p = p.Next
		}
		p.Next = node
	}
	list.length++
}

func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > list.length {
		return
	}

	if list.length == 0 {
		list.head = &ListNode{Val: val}
	} else {
		if index == 0 {
			list.AddAtHead(val)
		} else {
			p := list.head
			counter := 0
			for index-1 > counter {
				p = p.Next
				counter++
			}
			pn := p.Next
			p.Next = &ListNode{Val: val, Next: pn}
		}
	}

	list.length++
}

func (list *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= list.length {
		return
	}

	if list.length == 1 {
		list.head = nil
	} else if list.length == index+1 {
		p := list.head
		for p.Next.Next != nil {
			p = p.Next
		}
		p.Next = nil
	} else {
		if index == 0 {
			list.head = list.head.Next
		} else {
			p := list.head
			counter := 0
			for index-1 > counter {
				p = p.Next
				counter++
			}
			p.Next = p.Next.Next
		}
	}

	list.length--
}
