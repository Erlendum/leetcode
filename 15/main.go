package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	m := make(map[[3]int]bool)
	sort.Ints(nums)

	for i := range nums {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				m[[3]int{nums[i], nums[j], nums[k]}] = true
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	for k := range m {
		res = append(res, []int{k[0], k[1], k[2]})
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
