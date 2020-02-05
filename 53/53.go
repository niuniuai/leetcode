package main

import "fmt"

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArray(nums []int, left, right int) int {
	if right == left {
		return nums[left]
	}
	mid := (left + right) / 2
	maxLeft := maxArray(nums, left, mid)
	maxRight := maxArray(nums, mid+1, right)
	leftSubSum := nums[mid]
	curSum := 0
	for mid >= left {
		curSum += nums[mid]
		leftSubSum = maxInt(leftSubSum, curSum)
		mid--
	}

	mid = (left+right)/2 + 1
	rightSubSum := nums[mid]
	curSum = 0
	for mid <= right {
		curSum += nums[mid]
		rightSubSum = maxInt(rightSubSum, curSum)
		mid++
	}
	midMax := leftSubSum + rightSubSum
	return maxInt(maxInt(maxLeft, maxRight), midMax)

}

func maxSubArray(nums []int) int {
	return maxArray(nums, 0, len(nums)-1)
}

func main() {
	var nums []int = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	rel:= maxSubArray(nums)
	fmt.Println(rel)
}
