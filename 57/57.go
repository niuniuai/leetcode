package main

import (
	"fmt"
	"sort"
)

var a = [][]int{{1, 3}, {6, 9}}
var b = []int{2, 5}

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals=append(intervals,newInterval)
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	temp := intervals[0]
	var result [][]int
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= temp[1] {
			if intervals[i][1] > temp[1] {
				temp[1] = intervals[i][1]
			}

		} else {
			copydata:=make([]int,2)
			copy(copydata,temp)
			result = append(result, copydata)
			temp[0] = intervals[i][0]
			temp[1] = intervals[i][1]
		}
		if i == len(intervals)-1 {
			result = append(result, temp)
		}
	}
	return result
}

func main() {
	fmt.Println(insert(a, b))
}
