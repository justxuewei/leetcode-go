package bytedance

func twoSum(nums []int, target int) []int {
	hmap := make(map[int]int, len(nums))
	for i, v := range nums {
		hmap[v] = i
	}
	for i, v := range nums {
		key := target - v
		if ai, ok := hmap[key]; ok {
			if i == ai { continue }
			return []int{i, ai}
		}
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	hmap := make(map[int]int, len(nums))
	for i, v := range nums {
		another := target - v
		if anotherIdx, ok := hmap[another]; ok {
			return []int{anotherIdx, i}
		}
		hmap[v] = i
	}
	return nil
}
