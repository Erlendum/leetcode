package main

import (
	"fmt"
	"math"
)

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for i := range nums {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = 0
		}
		m[nums[i]]++
	}
	for k := range m {
		if m[k] == 1 {
			return k
		}
	}
	return math.MaxInt
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}
