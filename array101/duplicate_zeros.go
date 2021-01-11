package array101

func duplicateZeros(arr []int) {
	right := len(arr) - 1

	numOfZeros := 0

	removeLastElement := false
	for i, v := range arr {
		if i >= right - numOfZeros {
			if right - i < numOfZeros {
				removeLastElement = true
			}
			break
		}
		if v == 0 {
			numOfZeros++
		}
	}

	if numOfZeros == 0 {
		return
	}

	i := right
	if !removeLastElement {
		arr[right] = arr[right-numOfZeros]
		i--
	}

	for ; i>0; i-- {
		if arr[i-numOfZeros] == 0 {
			arr[i] = 0
			arr[i-1] = 0
			numOfZeros--
			i--
		} else {
			arr[i] = arr[i-numOfZeros]
		}
	}
}
