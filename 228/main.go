package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	res := make([]string, 0)

	var start, stop int

	nums = append(nums, nums[len(nums)-1])
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] == 1 {
			stop++
		} else {
			var s string
			if start < stop {
				startString := strconv.Itoa(nums[start])
				stopString := strconv.Itoa(nums[stop])
				s += startString + "->" + stopString
			} else {
				s += strconv.Itoa(nums[start])
			}
			res = append(res, s)
			start = i + 1
			stop = i + 1
		}
	}
	return res
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums))
}
