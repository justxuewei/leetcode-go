package bytedance

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	for i, a := range nums {
		if i > 0 && nums[i-1] == nums[i] { continue }
		target, j, k := -a, i+1, len(nums)-1
		for j < k {
			if nums[j] + nums[k] < target {
				j++
			} else if nums[j] + nums[k] > target {
				k--
			} else {
				ret = append(ret, []int{a, nums[j], nums[k]})
				for ; j < k && nums[j] == nums[j+1]; j++ {}
				for ; j < k && nums[k] == nums[k-1]; k-- {}
				j, k = j+1, k-1
			}
		}
	}
	return ret
}
