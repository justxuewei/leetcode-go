package jzoffer1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	numsString := make(Nums, len(nums))
	for i := range nums {
		numsString[i] = fmt.Sprintf("%d", nums[i])
	}

	sort.Sort(numsString)

	return strings.Join(numsString, "")
}

type Nums []string

func (n Nums) Len() int {
	return len(n)
}

func (n Nums) Less(i, j int) bool {
	x, _ := strconv.Atoi(n[i] + n[j])
	y, _ := strconv.Atoi(n[j] + n[i])
	return x < y
}

func (n Nums) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
