package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	start := intervals[0][0]
	stop := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= stop {
			if intervals[i][1] >= stop {
				stop = intervals[i][1]
			}
		} else {
			res = append(res, []int{start, stop})
			start = intervals[i][0]
			stop = intervals[i][1]
		}
	}
	res = append(res, []int{start, stop})
	return res
}

func main() {
	intervals := [][]int{{1, 4}, {2, 3}}
	fmt.Println(merge(intervals))
}
