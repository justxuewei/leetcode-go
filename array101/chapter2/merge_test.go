package array101

import "testing"

func TestMerge(t *testing.T) {
	nums1 := []int{2, 0}
	n := 1
	nums2 := []int{1}
	m := 1

	merge(nums1, n, nums2, m)
	solution2(nums1, n, nums2, m)
	t.Log(nums1)
}
