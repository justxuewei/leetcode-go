package jzoffer

import "sort"

func isStraight(nums []int) bool {
	sort.Ints(nums)
	any := 0
	for i:=0; i<len(nums)-1;i++ {
		if nums[i] == 0 {
			any++
		} else {
			if nums[i] == nums[i+1] {
				return false
			}
			any -= nums[i+1]-nums[i]-1
		}
		if any < 0 {
			return false
		}
	}
	return true
}
