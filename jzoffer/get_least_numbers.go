package jzoffer

import "sort"

// sort
func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 {
		return nil
	}
	sort.Ints(arr)
	return arr[:k]
}

// heap
func getLeastNumbers1(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}
	maxHeap := maxHeapInit(arr[:k])
	if len(arr) == k {
		return maxHeap
	}
	for i := k; i < len(arr); i++ {
		if maxHeap[0] > arr[i] {
			maxHeap[0] = arr[i]
			keepMaxHeapFromTopToBottom(maxHeap, 0)
		}
	}
	return maxHeap
}

func maxHeapInit(elements []int) []int {
	heap := make([]int, len(elements))
	// deep copy
	for i, v := range elements {
		heap[i] = v
	}
	for p := parent(len(heap) - 1); p >= 0; p-- {
		keepMaxHeapFromTopToBottom(heap, p)
	}
	return heap
}

func parent(i int) int {
	return ((i + 1) / 2) - 1
}

func children(i int) (int, int) {
	return 2*(i+1) - 1, 2 * (i + 1)
}

func maxChild(heap []int, p int) int {
	li, ri := children(p)
	if li >= len(heap) {
		return -1
	}
	if ri >= len(heap) {
		return li
	}
	if heap[li] > heap[ri] {
		return li
	} else {
		return ri
	}
}

func keepMaxHeapFromTopToBottom(heap []int, p int) {
	// c is an index of max child, sp is an index of subparent
	for c, sp := maxChild(heap, p), p; c != -1; sp, c = c, maxChild(heap, c) {
		if heap[sp] >= heap[c] {
			break
		}
		heap[sp], heap[c] = heap[c], heap[sp]
	}
}

func getLeastNumbers2(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}
	quicksort(arr, 0, len(arr)-1, k)
	return arr[:k]
}

func quicksort(arr []int, i, j, k int) {
	if i < 0 || j >= len(arr) || j-i < 1 {
		return
	}
	a, b := i+1, j
	for a <= b {
		if arr[b] < arr[i] {
			arr[a], arr[b] = arr[b], arr[a]
			a++
			continue
		}
		b--
	}
	arr[i], arr[a-1] = arr[a-1], arr[i]
	if k == a-1 {
		return
	} else if k > a-1 {
		quicksort(arr, a, j, k)
	} else {
		quicksort(arr, i, a-2, k)
	}
}
