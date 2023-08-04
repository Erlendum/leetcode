package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	flag := false
	for left != right && !flag {
		mid := (left + right) / 2
		if nums[mid] > target {
			flag = right == mid
			right = mid
		} else {
			flag = left == mid
			left = mid
		}
	}
	if nums[left] < target && nums[right] >= target {
		return right
	} else if nums[right] < target {
		return right + 1
	}
	return left
}

func main() {
	nums := []int{1, 3}
	target := 3
	fmt.Println(searchInsert(nums, target))
}
