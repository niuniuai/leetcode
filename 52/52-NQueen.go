package main

import (
	"fmt"
)

var (
	result int = 0
	daig1  = make(map[int]int) //记录主对角线值
	daig2  = make(map[int]int) //记录从对角线值
	rows   = make(map[int]int) //记录某行的某列是否被占用
)

func placeQueen(row, col int) {
	rows[col] = 1
	daig1[row-col] = 1
	daig2[row+col] = 1

}

func removeQueen(row, col int) {
	rows[col] = 0
	daig1[row-col] = 0
	daig2[row+col] = 0
}

func isAttack(row, col int) bool {
	res := rows[col] + daig1[row-col] + daig2[row+col]
	if res == 0 {
		return false
	}
	return true
}

func addSolution() {
	result++
}

func backtrack(row, n int) {
	for col := 0; col < n; col++ {
		if !isAttack(row, col) {
			placeQueen(row, col)
			if row == n-1 {
				addSolution()
			} else {
				backtrack(row+1, n)
			}
			removeQueen(row, col)
		}
	}
}

func totalNQueens(n int) int {
    backtrack(0, n)
	return result
}

func main() {
	fmt.Println(totalNQueens(1))
}
