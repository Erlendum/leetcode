package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	var total, avail, index int

	for i := range gas {
		total += gas[i] - cost[i]
		avail += gas[i] - cost[i]
		if avail < 0 {
			avail = 0
			index = i + 1
		}
	}

	if total < 0 {
		return -1
	}
	return index
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}
