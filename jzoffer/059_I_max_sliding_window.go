package jzoffer

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < 2 {
		return nums
	}
	monotonicQueue := make([]int, 1)
	ret := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		// remove head element if it is not in the sliding window
		if i-k >= 0 && nums[i-k] == monotonicQueue[0] {
			monotonicQueue = monotonicQueue[1:]
		}
		for j := 0; j < len(monotonicQueue); j++ {
			if nums[i] > monotonicQueue[j] {
				monotonicQueue = monotonicQueue[:j]
				break
			}
		}
		monotonicQueue = append(monotonicQueue, nums[i])
		if i >= k-1 {
			ret = append(ret, monotonicQueue[0])
		}
	}
	return ret
}
