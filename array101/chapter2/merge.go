package array101

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 { return }
	if m == 0 {
		for i, v := range nums2 {
			nums1[i] = v
		}
	}

	tmp := make([]int, len(nums1))

	var left, right []int
	var leftLen, rightLen int

	if n < m {
		left = nums2
		leftLen = n
		right = nums1
		rightLen = m
	} else {
		left = nums1
		leftLen = m
		right = nums2
		rightLen = n
	}

	i := 0
	j := 0
	for ; i<leftLen; i++ {
		cur := left[i]
		for j < rightLen && cur > right[j] {
			tmp[i+j] = right[j]
			j++
		}
		tmp[i+j] = left[i]
	}

	for ; j<rightLen; j++ {
		tmp[i+j] = right[j]
	}

	for i, v := range tmp {
		nums1[i] = v
	}
}

// ref: https://leetcode.com/problems/merge-sorted-array/discuss/197002/Go-solution
func solution2(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}
