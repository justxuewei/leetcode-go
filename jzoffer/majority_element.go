package jzoffer

// https://leetcode.com/problems/majority-element/
// https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
func majorityElement(nums []int) int {
	res, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			res, count = nums[i], 1
		} else {
			if res == nums[i] {
				count++
			} else {
				count--
			}
		}
	}
	return res
}
