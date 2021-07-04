package dp

import "math"

func jump(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 0
	for i:=1; i<len(nums); i++ {
		dp[i] = math.MaxInt32
	}

	for i:=0; i<len(nums)-1; i++ {
		end := i+1+nums[i]
		if end >= len(nums) {
			end = len(nums)
		}
		for j:=i+1; j<end; j++ {
			if dp[j] > dp[i]+1 {
				dp[j] = dp[i]+1
			}
		}
	}

	return dp[len(nums)-1]
}
