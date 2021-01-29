package chapter5

// runtime beats 26.54%
// memory usage beats 74.88%
func replaceElements(arr []int) []int {
	if len(arr) < 1 { return []int{} }

	for i:=0; i<len(arr)-1; i++ {
		max := arr[i+1]
		if i+2 < len(arr) {
			for j:=i+2; j<len(arr);j++ {
				if max < arr[j] {
					max = arr[j]
				}
			}
		}
		arr[i] = max
	}

	arr[len(arr)-1] = -1

	return arr
}

// runtime beats 99.05%
// memory usage beats 20.38%
func mySolution2(arr []int) []int {
	if len(arr) < 1 { return []int{} }

	maxVal := 0
	maxIndex := 0

	for i:=0; i<len(arr)-1; i++ {
		// max value should be updated
		if maxIndex == i {
			maxVal = arr[i+1]
			maxIndex = i+1
			// there is no more elements in array
			if i+2 < len(arr) {
				for j:=i+2; j<len(arr); j++ {
					if maxVal < arr[j] {
						maxVal = arr[j]
						maxIndex = j
					}
				}
			}
		}
		arr[i] = maxVal
	}

	// the last element is always -1
	arr[len(arr)-1] = -1

	return arr
}
