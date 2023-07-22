package main

import "fmt"

func majorityElement(nums []int) int {
	myMap := make(map[int]int, 0)

	for i := range nums {
		_, ok := myMap[nums[i]]
		if !ok {
			myMap[nums[i]] = 1
		} else {
			myMap[nums[i]]++
		}
	}
	for i := range myMap {
		if myMap[i] > len(nums)/2 {
			return i
		}
	}
	return 0
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}

	fmt.Println(majorityElement(nums))
}
