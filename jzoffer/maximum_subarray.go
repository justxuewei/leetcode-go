package jzoffer

import "math"

// greedy algorithm
func maxSubArray(nums []int) int {
	ret, cnt := math.MinInt32, 0

	for _, v := range nums {
		if cnt <= 0 {
			cnt = v
		} else {
			cnt += v
		}

		if ret < cnt {
			ret = cnt
		}
	}

	return ret
}

// dynamic programming
func maxSubArray1(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] = nums[i-1] + nums[i]
		}
	}

	ret := nums[0]
	for _, v := range nums {
		if ret < v {
			ret = v
		}
	}

	return ret
}

// divide and conquer
func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return math.MinInt32
	}
	if len(nums) == 1 {
		return nums[0]
	}

	mid := (len(nums) - 1) / 2
	// calculate maximum value of left subarray
	maxLeft, maxRight := 0, 0
	tmp := 0
	if mid > 0 {
		for i := mid - 1; i >= 0; i-- {
			tmp += nums[i]
			if tmp > maxLeft {
				maxLeft = tmp
			}
		}
	}
	tmp = 0

	if mid < len(nums)-1 {
		for i := mid + 1; i < len(nums); i++ {
			tmp += nums[i]
			if tmp > maxRight {
				maxRight = tmp
			}
		}
	}

	maxMid := maxLeft + maxRight + nums[mid]

	maxLeftMid, maxRightMid := maxSubArray(nums[:mid]), maxSubArray(nums[mid+1:])

	if maxMid >= maxRightMid && maxMid >= maxLeftMid {
		return maxMid
	}
	if maxLeftMid >= maxMid && maxLeftMid >= maxRightMid {
		return maxLeftMid
	}
	return maxRightMid
}
