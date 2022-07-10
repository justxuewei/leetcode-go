package problems

import (
	"math"
	"sort"
)

func getStrongest(arr []int, k int) []int {
	if len(arr) <= k {
		return arr
	}
	sort.Ints(arr)
	m := arr[(len(arr)-1)/2]

	var (
		l   = 0
		r   = len(arr) - 1
		ret = make([]int, 0, k)
	)

	for len(ret) < k && l < r {
		if math.Abs(float64(arr[l]-m)) <= math.Abs(float64(arr[r]-m)) {
			ret = append(ret, arr[r])
			r--
		} else {
			ret = append(ret, arr[l])
			l++
		}
	}

	return ret
}
