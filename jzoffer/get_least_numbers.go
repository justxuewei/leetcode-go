package jzoffer

import "sort"

// sort
func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 { return nil }
	sort.Ints(arr)
	return arr[:k]
}

// heap
//func getLeastNumbers1(arr []int, k int) []int {
//	if len(arr) == 0 { return nil }
//	heap := make(int[], k)
//
//}
