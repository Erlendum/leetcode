package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	n1 := make([]int, len(nums))
	for i := range n1 {
		n1[i] = 1
	}
	n2 := make([]int, len(nums))
	for i := range n2 {
		n2[i] = 1
	}

	prod := nums[0]
	for i := 1; i < len(nums); i++ {
		n1[i] = prod
		prod *= nums[i]
	}

	prod = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		n2[i] = prod
		prod *= nums[i]
	}

	output := make([]int, len(nums))
	for i := range output {
		output[i] = n1[i] * n2[i]
	}
	return output
}

func main() {
	nums := []int{-1, 1, 0, -3, 3}

	fmt.Println(productExceptSelf(nums))
}
