package bytedance

// LeetCode: 74
// Runtime: 0ms
// Memory usage: 2.8MB
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 { return false }
	mid := len(matrix)-1
	start, end := startAndEndElementsInRow(matrix[mid], target)
	if target < start {
		return searchMatrix(matrix[:mid], target)
	} else if end < target {
		return searchMatrix(matrix[mid+1:], target)
	} else {
		for _, v := range matrix[mid] {
			if v == target {
				return true
			}
		}
		return false
	}
}

func startAndEndElementsInRow(row []int, element int) (int, int) {
	if len(row) < 1 { return 0, 0 }
	return row[0], row[len(row)-1]
}

func searchMatrix1(matrix [][]int, target int) bool {
	if len(matrix) < 1 { return false }
	startIdx, endIdx := 0, len(matrix)-1
	var midIdx int
	for startIdx <= endIdx {
		midIdx = (startIdx + endIdx)/2
		startNum, endNum := startAndEndElementsInRow(matrix[midIdx], target)
		if target < startNum {
			endIdx = midIdx - 1
		} else if target > endNum {
			startIdx = midIdx + 1
		} else {
			for _, v := range matrix[midIdx] {
				if v == target {
					return true
				}
			}
			return false
		}
	}

	return false
}
