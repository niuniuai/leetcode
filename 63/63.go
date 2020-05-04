package main

import (
	"fmt"
)
var a=[][]int{{0,0,0},{0,1,0},{0,0,0}}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m:=len(obstacleGrid)
	if m==0{
		return 0
	}
	n:=len(obstacleGrid[0])
	matrix := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		matrix[i] = make([]int, n+1)
	}
	matrix[1][1]=1
	if obstacleGrid[0][0]==1{
		matrix[1][1]=0
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				continue
			}
			if obstacleGrid[i-1][j-1] == 1 {
				continue
			}
			matrix[i][j] = matrix[i][j-1] + matrix[i-1][j]
		}
	}
	return matrix[m][n]
}

func main() {
	fmt.Println(uniquePathsWithObstacles(a))
}
