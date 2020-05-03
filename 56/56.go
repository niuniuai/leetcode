package main

import (
	"fmt"
	"sort"
)

var a = [][]int{{1, 4}, {2, 3}}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
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
			// copy := []int{0, 0}
			// copy[0] = temp[0]
			// copy[1] = temp[1]
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
	fmt.Println(merge(a))
}
