package main

import (
	"strconv"
)

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
