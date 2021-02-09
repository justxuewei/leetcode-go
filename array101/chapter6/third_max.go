package chapter6

import "math"

func thirdMax(nums []int) int {

	topArr := make([]int, 3)

	topArr[0] = nums[0]
	length := 1
	startIndex := 0

	for i:=1; i<len(nums); i++ {
		if length == 3 {
			startIndex = i
			break
		}
		flag := false
		for j:=0; j<length; j++ {
			if nums[i] == topArr[j] {
				flag = true
				break
			}
		}
		if !flag {
			topArr[length] = nums[i]
			length++
		}
	}

	quickSort(topArr[:length])
	if length < 3 { return topArr[length-1] }

	for i:=startIndex; i<len(nums); i++ {
		switch {
		case nums[i] > topArr[0] && nums[i] < topArr[1]:
			topArr[0] = nums[i]
		case nums[i] > topArr[1] && nums[i] < topArr[2]:
			topArr[0] = topArr[1]
			topArr[1] = nums[i]
		case nums[i] > topArr[2]:
			topArr[0] = topArr[1]
			topArr[1] = topArr[2]
			topArr[2] = nums[i]
		}
	}

	return topArr[0]
}

// ref: https://books.halfrost.com/leetcode/ChapterFour/0414.Third-Maximum-Number/
func goBookSolution(nums []int) int {
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
	for _, v := range nums {
		if v > a {
			c = b
			b = a
			a = v
		} else if v < a && v > b {
			c = b
			b = v
		} else if v < b && v > c {
			c = v
		}
	}
	if c == math.MinInt64 {
		return a
	}
	return c
}
>>>>>>> c944218be1567dc33e48ab30dfb53785e2aa5038
