package main

import "fmt"

func rShift(nums []int, pos int, step int) {
	for i := len(nums) - 1; i >= pos; i-- {
		if i-step >= 0 {
			nums[i] = nums[i-step]
		} else {
			nums[i] = -1
		}
	}
}

func rotate(nums []int, k int) {
	k %= len(nums)
	save := make([]int, k)
	copy(save, nums[len(nums)-k:])
	rShift(nums, 0, k)
	for i := 0; i < k; i++ {
		nums[i] = save[i]
	}
}

func main() {
	nums := []int{-1}
	k := 2
	rotate(nums, k)
	fmt.Println(nums)
}
