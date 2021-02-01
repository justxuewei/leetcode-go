package chapter6

import (
	"fmt"
	"math/rand"
)

func heightChecker(heights []int) int {
	heightsCpy := make([]int, len(heights))
	copy(heightsCpy, heights)
	quickSort(heights)

	fmt.Print(heights)
	moveSize := 0
	for i:=0; i<len(heights); i++ {
		if heights[i] != heightsCpy[i] { moveSize++ }
	}

	return moveSize
}

func quickSort(arr []int) {
	if len(arr) < 2 { return }

	randInt := rand.Intn(100) % len(arr)
	arr[0], arr[randInt] = arr[randInt], arr[0]
	pivot, index := 0, 1

	for i:=1; i<len(arr); i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}

	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]

	quickSort(arr[:index-1])
	quickSort(arr[index:])
}
