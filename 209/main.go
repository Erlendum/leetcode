package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	j := 0
	sum := 0
	res := len(nums) + 1
	for i := range nums {
		sum += nums[i]
		if sum >= target {
			for sum >= target {
				sum -= nums[j]
				j++
			}
			if i-j+2 < res {
				res = i - j + 2
			}
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	fmt.Println(minSubArrayLen(target, nums))
}
