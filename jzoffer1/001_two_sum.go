package jzoffer1

func twoSum(nums []int, target int) []int {
	set := make(map[int]int)

	for i, v := range nums {
		if v > target {
			continue
		}

		if ai, ok := set[target-v]; ok {
			return []int{i, ai}
		}
		set[v] = i
	}

	return nil
}
