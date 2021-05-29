package jzoffer

// TODO: Learn how go implements a heap container?(ref:https://www.cnblogs.com/huxianglin/p/6925119.html)

type Heap []int

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Pop() interface{} {
	ret := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return ret
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

type MinHeap struct {
	Heap
}

func (h *MinHeap) Less(i, j int) bool {
	return h.Heap[i] < h.Heap[j]
}

type MaxHeap struct {
	Heap
}

func (h *MaxHeap) Less(i, j int) bool {
	return h.Heap[i] > h.Heap[j]
}

type MedianFinder struct {
	maxh *MaxHeap
	minh *MinHeap
}

func NewMedianFinder() MedianFinder {
	return MedianFinder{
		maxh: &MaxHeap{make(Heap, 0)},
		minh: &MinHeap{make(Heap, 0)},
	}
}

func (f *MedianFinder) AddNum(num int) {
	if f.minh.Len() == f.maxh.Len() {
		f.maxh.Push(num)
		f.minh.Push(f.maxh.Pop())
	} else {
		f.minh.Push(num)
		f.maxh.Push(f.minh.Pop())
	}
}

func (f *MedianFinder) FindMedian() float64 {
	if f.minh.Len() == f.maxh.Len() {
		return (float64(f.minh.Heap[0]) + float64(f.maxh.Heap[0])) / 2.0
	}
	return float64(f.minh.Heap[0])
}
