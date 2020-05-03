package main

import (
	"fmt"
)

func generateMatrix(n int) [][]int {
	level := (n + 1) / 2
	startX, startY := 0, 0
	endX, endY := n-1, n-1
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	count := 1
	for l := 1; l <= level; l++ {
		for j := startY; j <= endY; j++ {
			result[startX][j] = count
			count += 1
		}
		for i := startX + 1; i <= endX; i++ {
			result[i][endY] = count
			count += 1
		}
		for j := endY - 1; j >= startY; j-- {
			result[endX][j] = count
			count += 1
		}
		for i := endX - 1; i > startX; i-- {
			result[i][startY] = count
			count += 1
		}
		startX++
		startY++
		endX--
		endY--
	}
	return result
}

func main() {
	fmt.Println(generateMatrix(0))
}
