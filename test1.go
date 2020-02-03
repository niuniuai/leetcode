package main

import (
	"fmt"
)

var (
	output [][]string
	daig1  = make(map[int]int) //记录主对角线值
	daig2  = make(map[int]int) //记录从对角线值
	rows   = make(map[int]int) //记录某行的某列是否被占用
	queens = make(map[int]int) //用于记录第几行的第几列可放置皇后
)

func placeQueen(row, col, n int) {
	rows[col] = 1
	daig1[row-col+2*n] = 1
	daig2[row+col] = 1
	queens[row] = col

}

func removeQueen(row, col, n int) {
	rows[col] = 0
	daig1[row-col+2*n] = 0
	daig2[row+col] = 0
	queens[row] = 0
}

func isAttack(row, col, n int) bool {
	res := rows[col] + daig1[row-col+2*n] + daig2[row+col]
	if res == 0 {
		return false
	}
	return true
}

func addSolution(n int) {
	var rel []string
	for i := 0; i < n; i++ {
		var str string
		for j := 0; j < queens[i]; j++ {
			str += "."
		}
		str += "Q"
		for j := queens[i]+1; j < n; j++ {
			str += "."
		}
		rel = append(rel, str)
	}
	output = append(output, rel)
}

func backtrack(row, n int) {
	for col := 0; col < n; col++ {
		if !isAttack(row, col, n) {
			placeQueen(row, col, n)
			if row == n-1 {
				addSolution(n)
			} else {
				backtrack(row+1, n)
			}
			removeQueen(row, col, n)
		}
	}
}

func solveNQueens(n int) [][]string {
	backtrack(0, n)
	return output
}

func main() {
	fmt.Println(solveNQueens(1))
}
