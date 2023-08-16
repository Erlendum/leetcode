package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	maxSum := math.MinInt
	curSum := 0
	for i := range nums {
		curSum += nums[i]
		if curSum > maxSum {
			maxSum = curSum
		}
		if curSum < 0 {
			curSum = 0
		}
	}
	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
