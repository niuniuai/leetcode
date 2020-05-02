package main

import (
	"fmt"
)

var a = []int{3,2,1,0,4}

func canJump(nums []int) bool {
	length := len(nums)
	max := 0
	for i := 0; i < length-1; i++ {
		if max < nums[i] {
			max = nums[i]
		}
		max--
		if max<0{
			return false
		}
	}
	if max<0{
		return false
	}
	return true
}

func main() {
	fmt.Println(canJump(a))
}
