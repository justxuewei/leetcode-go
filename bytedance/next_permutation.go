package bytedance

func nextPermutation(nums []int) {
	prepeek := -1
	for i:=len(nums)-1; i>0; i-- {
		if nums[i] > nums[i-1] {
			prepeek = i-1
			break
		}
	}

	if prepeek == -1 {
		reverse(nums)
		return
	}

	for i:=len(nums)-1; i>prepeek; i-- {
		if nums[i] > nums[prepeek] {
			nums[prepeek], nums[i] = nums[i], nums[prepeek]
			break
		}
	}
	quickSort(nums[prepeek+1:])
}

func reverse(arr []int) {
	for i:=0; i<len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
