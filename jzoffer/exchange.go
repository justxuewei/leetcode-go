package jzoffer

// https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
func exchange(nums []int) []int {
	if len(nums) < 2 { return nums }

	left, right := 0, len(nums)-1
	for {
		for left < right && nums[left] % 2 != 0 {
			left++
		}

		if left >= right { break }

		for left < right && nums[right] % 2 == 0 {
			right--
		}

		if left >= right { break }

		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}

func exchange1(nums []int) []int {
	if len(nums) < 2 { return nums }

	for slow, fast := 0, 0; fast < len(nums); fast++ {
		if nums[fast] & 1 == 1 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}

	return nums
}
