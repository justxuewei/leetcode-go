package chapter6

func findDisappearedNumbers(nums []int) []int {

	for _, v := range nums {
		tmp := v
		for nums[tmp-1] != tmp {
			nums[tmp-1], tmp = tmp, nums[tmp-1]
		}
	}

	ret := make([]int, 0)
	for i, v := range nums {
		if v != i+1 {
			ret = append(ret, i+1)
		}
	}

	return ret
}
