package dp

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	ret := nums[0]
	for i:=1; i<len(nums); i++ {
		if nums[i-1] + nums[i] > nums[i] {
			nums[i] += nums[i-1]
		}
		if ret < nums[i] {
			ret = nums[i]
		}
	}

	return ret
}
