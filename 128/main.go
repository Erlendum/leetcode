package main

import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	sort.Ints(nums)

	maxLen := 0
	curLen := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			if nums[i]-nums[i-1] == 1 {
				curLen++
			} else {
				if curLen > maxLen {
					maxLen = curLen
				}
				i -= curLen - 1
				curLen = 1
			}
		}
	}

	if curLen > maxLen {
		maxLen = curLen
	}

	return maxLen
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
