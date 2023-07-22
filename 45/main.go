package main

import (
	"fmt"
	"math"
)

func myMin(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func myJump(nums []int, index int, m *int, c int, memo []int) bool {
	if index >= len(nums) {
		return false
	}

	if memo[index] != 0 && c >= memo[index] {
		return false
	}
	memo[index] = c
	if c >= *m {
		return false
	}

	if index == len(nums)-1 {
		*m = myMin(*m, c)
		return true
	}
	if nums[index] == 0 {
		return false
	}

	for i := nums[index]; i > 0; i-- {
		if i < len(nums) {
			myJump(nums, index+i, m, c+1, memo)
		}
	}

	return false
}
func jump(nums []int) int {
	memo := make([]int, len(nums))

	var m int
	m = math.MaxInt

	myJump(nums, 0, &m, 0, memo)
	return m
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Print(jump(nums))
}
