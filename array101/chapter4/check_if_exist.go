package chapter4

func checkIfExist(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j]*2 == arr[i] || arr[i]*2 == arr[j] {
				return true
			}
		}
	}
	return false
}
