package main

func searchMatrix(matrix [][]int, target int) bool {
	//do two binary searchesno?
	rows := len(matrix)
	columns := len(matrix[0])

	top := 0
	bot := rows - 1

	for top <= bot {
		row := (top + bot) / 2
		if target < matrix[row][0] {
			bot = row - 1
		} else if target > matrix[row][columns-1] {
			top = row + 1
		} else {
			break
		}
	}

	row := (top + bot) / 2

	left := 0
	right := columns - 1

	for left <= right {
		col := (left + right) / 2
		if target == matrix[row][col] {
			return true
		} else if target > matrix[row][col] {
			left = col + 1
		} else {
			right = col - 1
		}
	}
	return false
}
