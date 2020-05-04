package main

import (
	"fmt"
)

func uniquePaths(m int, n int) int {
	matrix := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		matrix[i] = make([]int, n+1)
	}
	matrix[1][1] = 1
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				continue
			}
			matrix[i][j] = matrix[i][j-1] + matrix[i-1][j]
		}
	}
	return matrix[m][n]
}

func main() {
	fmt.Println(uniquePaths(3, 2))
}
