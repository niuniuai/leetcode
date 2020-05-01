package main

import (
	"fmt"
)

var a = [][]int{
	{1, 2, 3, 4},
	{5, 6, 7, 8},
	{9, 10, 11, 12},
}

func spiralOrder(matrix [][]int) []int {
	endX := len(matrix) - 1
	if endX == -1 {
		return nil
	}
	endY := len(matrix[0]) - 1
	startX, startY := 0, 0
	var result []int
	for startX <= endX && startY <= endY {
		for j := startY; j <= endY; j++ {
			result = append(result, matrix[startX][j])
		}
		for i := startX + 1; i <= endX; i++ {
			result = append(result, matrix[i][endY])
		}
		if startX < endX && startY < endY {
			for j := endY - 1; j >= startY; j-- {
				result = append(result, matrix[endX][j])
			}
			for i := endX - 1; i > startX; i-- {
				result = append(result, matrix[i][startY])
			}

		}
		startX++
		startY++
		endX--
		endY--
	}
	return result
}

func main() {
	fmt.Println(spiralOrder(a))
}
