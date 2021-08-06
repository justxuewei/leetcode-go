package jzoffer1

import (
	"container/heap"
	"math"
)

type MedianFinder struct {
	left *IntMaxHeap
	right *IntMinHeap
}

func NewMedianFinder() MedianFinder {
	left, right := &IntMaxHeap{}, &IntMinHeap{}
	heap.Init(left)
	heap.Init(right)
	return MedianFinder{
		left: left,
		right: right,
	}
}

func (f *MedianFinder) AddNum(num int)  {
	if f.left.Len() == 0 || (*f.left)[0] >= num {
		heap.Push(f.left, num)
	} else {
		heap.Push(f.right, num)
	}

	if math.Abs(float64(f.left.Len() - f.right.Len())) > 1 {
		if f.left.Len() > f.right.Len() {
			heap.Push(f.right, heap.Pop(f.left))
		} else {
			heap.Push(f.left, heap.Pop(f.right))
		}
	}
}


func (f *MedianFinder) FindMedian() float64 {
	if f.left.Len() == f.right.Len() {
		return (float64((*f.left)[0]) + float64((*f.right)[0])) / 2.0
	} else if f.left.Len() > f.right.Len() {
		return float64((*f.left)[0])
	} else {
		return float64((*f.right)[0])
	}
}

type IntMinHeap []int

func (h IntMinHeap) Len() int {
	return len(h)
}

func (h IntMinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntMinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0: n-1]
	return x
}

type IntMaxHeap []int

func (h IntMaxHeap) Len() int {
	return len(h)
}

func (h IntMaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntMaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntMaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0: n-1]
	return x
}