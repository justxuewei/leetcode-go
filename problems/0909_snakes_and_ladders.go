package problems

func snakesAndLadders(board [][]int) int {
	var (
		stepZero = &ListNode{
			Val:  1,
			Next: nil,
		}
		head = &ListNode{
			Val:  0,
			Next: stepZero,
		}
		stepLast = stepZero
		current  *ListNode
		next     *ListNode
		nextVal  int
		tail     = stepZero
		step     = 1
		i        = 0
		n2       = len(board) * len(board)
		row, col int
		numSet   = make(map[int]int8)
		ok       bool
	)

	for head.Next != nil {
		current = head.Next

		for i = 1; i <= 6; i++ {
			nextVal = current.Val + i
			
			row, col = getCoordination(len(board), nextVal)
			if board[row][col] != -1 {
				nextVal = board[row][col]
			}

			if _, ok = numSet[nextVal]; ok {
				continue
			} else {
				numSet[nextVal] = 0
			}
			if nextVal >= n2 {
				return step
			}

			next = &ListNode{Val: nextVal}

			tail.Next = next
			tail = next
		}

		if current.Val == stepLast.Val {
			stepLast = tail
			step++
		}

		head.Next = current.Next
	}

	return -1

}

func getCoordination(n, num int) (row int, col int) {
	row = n - (num-1)/n - 1
	col = num%n - 1
	if col < 0 {
		col = n - 1
	}
	if row%2 == n%2 {
		col = (n - 1) - col
	}
	return
}
