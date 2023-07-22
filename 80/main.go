package main

import "fmt"

func lShift(nums []int, pos int, step int) {
	for i := pos; i < len(nums); i++ {
		if i+step < len(nums) {
			nums[i] = nums[i+step]
		} else {
			nums[i] = -101
		}
	}
}

func seqLen(nums []int, pos int) int {
	c := 0
	for i := pos + 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			return c
		}
		c++
	}
	return c
}

func removeDuplicates(nums []int) int {
	c := 0
	for i := 0; i < len(nums) && nums[i] != -101; i++ {
		step := seqLen(nums, i)
		if step == 1 {
			continue
		}
		step--
		if step > 0 {
			lShift(nums, i, step)
			c += step
		}
	}
	return len(nums) - c
}

func main() {
	nums := []int{0, 0, 1, 1, 2, 2, 3, 4, 5, 6, 6, 6, 6, 6, 7}

	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
