package main

import (
	"fmt"
	"math"
)

func maxSubarraySumCircular(nums []int) int {
	sum := 0
	maxSum := math.MinInt
	curSum := 0
	for i := range nums {
		sum += nums[i]
		curSum += nums[i]
		if curSum > maxSum {
			maxSum = curSum
		}
		if curSum < 0 {
			curSum = 0
		}
	}
	minSum := math.MaxInt
	curSum = 0
	for i := range nums {
		curSum += nums[i]
		if curSum < minSum {
			minSum = curSum
		}
		if curSum > 0 {
			curSum = 0
		}
	}

	if sum == minSum {
		return maxSum
	}

	if maxSum > sum-minSum {
		return maxSum
	}
	return sum - minSum
}

func main() {
	nums := []int{5, -3, 5}
	fmt.Println(maxSubarraySumCircular(nums))
}
