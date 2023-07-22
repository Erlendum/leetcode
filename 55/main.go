package main

import "fmt"

func jump(nums []int, index int, memo []bool) bool {
	if index >= len(nums) {
		return false
	}
	if memo[index] {
		return false
	}

	memo[index] = true

	if index == len(nums)-1 {
		return true
	}
	if nums[index] == 0 {
		return false
	}
	if nums[index] >= len(nums)-1 {
		return true
	}
	for i := nums[index]; i > 0; i-- {
		if i < len(nums) && jump(nums, index+i, memo) {
			return true
		}
	}

	return false
}
func canJump(nums []int) bool {
	memo := make([]bool, len(nums))
	for i := range memo {
		memo[i] = false
	}
	if jump(nums, 0, memo) {
		return true
	}
	return false
}

func main() {
	nums := []int{2, 1, 2, 1, 2, 1}
	fmt.Print(canJump(nums))
}
