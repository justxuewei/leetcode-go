package jzoffer

func missingNumber(nums []int) int {
	for i:=1; i<len(nums); i++ {
		if nums[i] - nums[i-1] > 1 {
			return nums[i]-1
		}
	}
	if nums[0] == 0 {
		return nums[len(nums)-1]+1
	}
	return 0
}
