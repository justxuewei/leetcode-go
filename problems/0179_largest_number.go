package problems

import (
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	numStrings := make([]string, len(nums))
	var numString string
	for i, num := range nums {
		numString = strconv.Itoa(num)
		numStrings[i] = numString
	}

	quickSort(numStrings)
	return strings.Join(numStrings, "")
}

func less(lhs, rhs string) bool {
	var (
		lhsIdx = len(lhs)
		rhsIdx = len(rhs)
	)

	for lhsIdx >= 0 && rhsIdx >= 0 {
		lhsIdx--
		rhsIdx--

		if lhs[lhsIdx] != rhs[rhsIdx] {
			return lhs[lhsIdx] < rhs[rhsIdx]
		}
	}

	if lhsIdx < 0 &&  {

	}

	return false
}

func quickSort(arr []string) {
	if len(arr) <= 1 {
		return
	}

	l, r := 1, len(arr)-1

	for l <= r {
		if less(arr[r], arr[0]) {
			r--
			continue
		}
		arr[l], arr[r] = arr[r], arr[l]
		l++
	}

	arr[0], arr[r] = arr[r], arr[0]
	quickSort(arr[:r])
	quickSort(arr[r+1:])
}
