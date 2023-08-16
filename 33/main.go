package main

import (
	"fmt"
	"math"
)

func minIndex(nums []int) int {
	minI := -1
	min := math.MaxInt

	for i := range nums {
		if nums[i] < min {
			min = nums[i]
			minI = i
		}
	}
	return minI
}

func search(nums []int, target int) int {
	shift := minIndex(nums)

	l := shift
	r := len(nums) + shift - 1

	for l <= r {
		mid := (l + r) / 2 % len(nums)
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = (l+r)/2 - 1
		} else {
			l = (l+r)/2 + 1
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3
	fmt.Println(search(nums, target))
}
