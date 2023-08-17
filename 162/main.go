package main

import "fmt"

func findPeakElement(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := (l + r) / 2
		if (mid == 0 || nums[mid] >= nums[mid-1]) && (mid == len(nums)-1 || nums[mid] >= nums[mid+1]) {
			return mid
		} else if nums[mid] <= nums[mid+1] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(findPeakElement(nums))
}
