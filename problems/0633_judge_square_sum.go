package problems

import (
	"math"
)

func judgeSquareSum(c int) bool {
	if c <= 2 {
		return true
	}
	left, right := 0, c
	var start, mid, end, result int
	for left < right {
		if int(math.Pow(float64(left), 2)+math.Pow(float64(left), 2)) == c {
			return true
		}

		start, end = left+1, right
		for start < end {
			mid = (start + end) / 2
			result = int(math.Pow(float64(left), 2) + math.Pow(float64(mid), 2))
			if result < c {
				if start == mid {
					break
				}
				start = mid
			} else if result > c {
				if end == mid {
					break
				}
				end = mid
			} else {
				return true
			}
		}
		left++
		right = end
	}

	return false
}

func judgeSquareSum1(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum > c {
			right--
		} else {
			left++
		}
	}
	return false
}
