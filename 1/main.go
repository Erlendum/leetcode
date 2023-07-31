package main

import "fmt"

func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int, 0)

	for i := range nums {
		if el, ok := myMap[nums[i]]; ok {
			return []int{el, i}
		}
		myMap[target-nums[i]] = i
	}
	return []int{}
}

func main() {
	nums := []int{3, 3}
	target := 6
	fmt.Println(twoSum(nums, target))
}
