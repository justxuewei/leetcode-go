package bytedance

func heapSort(nums []int) {
	// build heap
	parent := parentIndex(len(nums)-1)
	for parent >= 0 {
		if isSwapped, child := swapMaxChild(nums, parent); isSwapped {
			newParent := child
			for newParent < len(nums) {
				if isSwapped, newChild := swapMaxChild(nums, newParent); isSwapped {
					newParent = newChild
				} else {
					break
				}
			}
		}
		parent = parent - 1
	}

	unsortedNums := nums
	for len(unsortedNums) > 0 {
		unsortedNums[0], unsortedNums[len(unsortedNums)-1] = unsortedNums[len(unsortedNums)-1], unsortedNums[0]
		unsortedNums = unsortedNums[:len(unsortedNums)-1]
		parent = 0
		for parent < len(unsortedNums) {
			if isSwapped, child := swapMaxChild(unsortedNums, parent); isSwapped {
				parent = child
			} else {
				break
			}
		}
	}
}

func parentIndex(i int) int {
	return ((i + 1) / 2) - 1
}

func childrenIndexes(i int) (int, int) {
	return 2 * (i + 1) - 1, (2 * (i + 1) + 1) - 1
}

func swapMaxChild(nums []int, parent int) (bool, int) {
	left, right := childrenIndexes(parent)
	var maxChild int
	if left >= len(nums) {
		return false, -1
	}
	if right >= len(nums) {
		maxChild = left
	} else {
		if nums[left] > nums[right] {
			maxChild = left
		} else {
			maxChild = right
		}
	}

	if nums[maxChild] > nums[parent] {
		nums[maxChild], nums[parent] = nums[parent], nums[maxChild]
		return true, maxChild
	}
	return false, -1
}


