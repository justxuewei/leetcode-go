package jzoffer1

func singleNumbers(nums []int) []int {
	if len(nums) <= 2 {
		return nums
	}

	xor := nums[0]
	for i:=1; i<len(nums); i++ {
		xor ^= nums[i]
	}

	flag := 0
	for (xor >> flag) & 1 == 0 {
		flag++
	}

	ret := make([]int, 2)
	for i:=0; i<len(nums); i++ {
		if ret[(nums[i] >> flag) & 1] == 0 {
			ret[(nums[i] >> flag)  & 1] = nums[i]
		} else {
			ret[(nums[i] >> flag)  & 1] ^= nums[i]
		}
	}

	return ret
}
