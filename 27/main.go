package main

import "fmt"

func lShift(nums []int, pos int) {
	for i := pos; i < len(nums); i++ {
		if i != len(nums)-1 {
			nums[i] = nums[i+1]
		} else {
			nums[i] = -1
		}
	}
}

func removeElement(nums []int, val int) int {
	c := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			lShift(nums, i)
			i--
			c++
		}
	}
	return len(nums) - c
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}

	fmt.Println(removeElement(nums, 2))
	fmt.Println(nums)
}
