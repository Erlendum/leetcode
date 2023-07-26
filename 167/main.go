package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	for index1 := range numbers {
		for index2 := index1 + 1; index2 < len(numbers) && numbers[index2]+numbers[index1] <= target; index2++ {
			if numbers[index2]+numbers[index1] == target {
				return []int{index1 + 1, index2 + 1}
			}
		}
	}
	return nil
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbers, target))
}
