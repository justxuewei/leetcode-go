package array101

func sortedSquares(nums []int) []int {
	arr := make([]int, len(nums))
	for i, v := range nums {
		arr[i] = v * v
	}



	return arr
}

func quickSort(nums []int, b int, e int) {
	if b > e { return }

	pivot := nums[b]
	i := b + 1
	j := e
	for i < j {
		for i < j {
			i++
			if nums[i] <= pivot { break }
		}

		for j > i {
			j--
			if nums[j] >= pivot { break }
		}
		if nums[i] > pivot && nums[j] < pivot {
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
		}
	}
	m := 0
	if nums[i] > pivot {
		m = i - 1
	} else {
		m = i
	}
	tmp := nums[m]
	nums[m] = nums[b]
	nums[b] = tmp
	quickSort(nums, b, m-1)
	quickSort(nums, m+1, e)
}
