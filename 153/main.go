package main

import "fmt"

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l := 0
	r := len(nums) - 1

	prev := nums[r]
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < prev {
			r = mid - 1
		} else {
			l = mid + 1
		}
		prev = nums[mid]
	}
	return nums[l]
}

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(nums))
}
