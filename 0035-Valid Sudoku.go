package main

import (
	"strconv"
)

/*This problem seems complicated but the constraints can be read as:
Make sure each number in each row, column and 3x3 grid is unique.
This solution stores the data as the key of a map with a specific format
If the number already exists then the sudoku board is invalid.
Note that this problem is a lot easier since you are only checking for if the current
positions of the numbers are valid but not accounting for caclulating empty squares.*/

func isValidSudoku(board [][]byte) bool {
	hashMap := make(map[string]struct{})
	for i, row := range board {
		for j, val := range row {
			current_val := string(val)
			if current_val == "." {
				continue
			}
			_, ok1 := hashMap[current_val+"row"+strconv.Itoa(i)]
			_, ok2 := hashMap[current_val+"column"+strconv.Itoa(j)]
			_, ok3 := hashMap[current_val+"grid"+strconv.Itoa(i/3)+"-"+strconv.Itoa(j/3)]
			if ok1 || ok2 || ok3 {
				return false
			} else {
				hashMap[current_val+"row"+strconv.Itoa(i)] = struct{}{}
				hashMap[current_val+"column"+strconv.Itoa(j)] = struct{}{}
				hashMap[current_val+"grid"+strconv.Itoa(i/3)+"-"+strconv.Itoa(j/3)] = struct{}{}
			}
		}
	}

	return true
}
