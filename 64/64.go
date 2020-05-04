package main

import (
	"fmt"
)

var a = [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}

func min(x, y int) int{
	if x < y {
		return x
	}
	return y
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	matrix[0][0]=grid[0][0]
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if i==0&&j==0{
				continue
			}
			if i==0{
				matrix[0][j]=grid[0][j]+matrix[0][j-1]
				continue
			}
			if j==0{
				matrix[i][0]=grid[i][0]+matrix[i-1][0]
				continue
			}
			matrix[i][j]=grid[i][j]+min(matrix[i-1][j],matrix[i][j-1])
		}
	}
	return matrix[m-1][n-1]
}

func main() {
	fmt.Println(minPathSum(a))
}
