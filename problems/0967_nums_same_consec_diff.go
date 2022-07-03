package problems

import (
	"fmt"
	"strconv"
)

func numsSameConsecDiff(n, k int) []int {
	ret := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var tmp int
	power := 1
	for i := 1; i < n; i++ {
		var current []int
		for _, retv := range ret {
			tmp = retv / power
			for j := 0; j < 10; j++ {
				if tmp-j == k || j-tmp == k {
					current = append(current, retv+j*power*10)
				}
			}
		}
		ret = current
		power *= 10
	}

	newRet := []int{}
	for _, retv := range ret {
		if retv/power != 0 {
			newRet = append(newRet, retv)
		}
	}

	return newRet
}

func numsSameConsecDiff1(n, k int) []int {
	var ret []int

	for i := 1; i < 10; i++ {
		dfs(fmt.Sprintf("%d", i), &ret, n, k)
	}

	return ret
}

func dfs(leftStr string, ret *[]int, n, k int) {
	if len(leftStr) == n {
		intV, _ := strconv.Atoi(leftStr)
		*ret = append(*ret, intV)
		return
	}

	lastDigit, _ := strconv.Atoi(string(leftStr[len(leftStr)-1]))
	bigger := lastDigit + k
	smaller := lastDigit - k
	if bigger < 10 {
		dfs(fmt.Sprintf("%s%d", leftStr, bigger), ret, n, k)
	}
	if k != 0 && smaller >= 0 {
		dfs(fmt.Sprintf("%s%d", leftStr, smaller), ret, n, k)
	}
}
