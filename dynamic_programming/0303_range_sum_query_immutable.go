package dynamic_programming

type NumArray struct {
	data []int
}

func NewNumArray(nums []int) NumArray {
	data := make([]int, 0, len(nums))
	data = append(data, nums[0])
	for i:=1; i<len(nums); i++ {
		data = append(data, data[i-1] + nums[i])
	}
	return NumArray{data: data}
}

func (a *NumArray) SumRange(left, right int) int {
	if left == 0 {
		return a.data[right]
	}
	return a.data[right] - a.data[left-1]
}
