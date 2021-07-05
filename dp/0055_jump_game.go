package dp

func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true

	var end int
	for i:=0; i<len(nums) && !dp[len(nums)-1]; i++ {
		if !dp[i] { continue }
		end = min(len(nums), i+1+nums[i])
		for j:=i+1; j<end; j++ {
			dp[j] = true
		}
	}

	return dp[len(nums)-1]
}