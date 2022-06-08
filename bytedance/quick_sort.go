package bytedance

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[0]
	curIdx := 1
	for i := len(arr) - 1; i >= curIdx; {
		if arr[i] < pivot {
			arr[i], arr[curIdx] = arr[curIdx], arr[i]
			curIdx++
			continue
		}
		i--
	}
	arr[0], arr[curIdx-1] = arr[curIdx-1], arr[0]
	quickSort(arr[:curIdx])
	quickSort(arr[curIdx:])
}
