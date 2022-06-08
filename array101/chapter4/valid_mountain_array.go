package chapter4

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	down := false
	if arr[0] >= arr[1] {
		return false
	}
	for i := 1; i < len(arr)-1; i++ {
		if !down && arr[i] >= arr[i+1] {
			down = true
		}
		if down && arr[i] <= arr[i+1] {
			return false
		}
	}
	if !down {
		return false
	}
	return true
}

// ref: https://leetcode.com/problems/valid-mountain-array/solution/
// standard solution has the same performance as mine
func standardSolution(arr []int) bool {
	i := 0

	// walk up
	for i+1 < len(arr) && arr[i] < arr[i+1] {
		i++
	}

	if i == 0 || i == len(arr)-1 {
		return false
	}

	// walk down
	for i+1 < len(arr) && arr[i] > arr[i+1] {
		i++
	}

	return i+1 == len(arr)
}
