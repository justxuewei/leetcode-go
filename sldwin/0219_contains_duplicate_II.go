package sldwin

func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 {
		return false
	}

	set := make(map[int]bool, k+1)

	for i:=0; i<len(nums); i++ {
		if _, ok := set[nums[i]]; ok {
			return true
		}

		if len(set) >= k {
			delete(set, nums[i-k])
		}

		set[nums[i]] = true
	}

	return false
}
