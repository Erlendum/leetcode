package main

import "fmt"

func binSearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
func searchRange(nums []int, target int) []int {
	start := binSearch(nums, target)
	stop := binSearch(nums, target+1) - 1
	if start < len(nums) && nums[start] == target {
		return []int{start, stop}
	}
	return []int{-1, -1}
}

func main() {
	nums := []int{5, 5, 5, 8, 8, 10}
	target := 10
	fmt.Println(searchRange(nums, target))
}
