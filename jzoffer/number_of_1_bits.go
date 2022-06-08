package jzoffer

// https://leetcode.com/problems/number-of-1-bits/
// https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
func hammingWeight(num uint32) int {
	nzeros := 0
	for num > 0 {
		if num&1 != num&0 {
			nzeros++
		}
		num >>= 1
	}
	return nzeros
}
