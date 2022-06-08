package array101

import "math/rand"

func sortedSquares(nums []int) []int {
	arr := make([]int, len(nums))
	for i, v := range nums {
		arr[i] = v * v
	}

	return quickSort(arr)
}

// func quickSort(nums []int, b int, e int) {
//	if b > e { return }
//
//	pivot := nums[b]
//	i := b + 1
//	j := e
//	for i < j {
//		for i < j {
//			i++
//			if nums[i] <= pivot { break }
//		}
//
//		for j > i {
//			j--
//			if nums[j] >= pivot { break }
//		}
//		if nums[i] > pivot && nums[j] < pivot {
//			tmp := nums[i]
//			nums[i] = nums[j]
//			nums[j] = tmp
//		}
//	}
//	m := 0
//	if nums[i] > pivot {
//		m = i - 1
//	} else {
//		m = i
//	}
//	tmp := nums[m]
//	nums[m] = nums[b]
//	nums[b] = tmp
//	quickSort(nums, b, m-1)
//	quickSort(nums, m+1, e)
// }

// refL: https://www.golangprograms.com/golang-program-for-implementation-of-quick-sort.html
func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Intn(100) % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}
