package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	p := citations[0]
	for p > 0 {
		i := 0
		for ; i < len(citations) && i < p; i++ {
			if citations[i] < p {
				break
			}
		}
		if i == p {
			return p
		}
		p--
	}
	return p
}

func main() {
	citations := []int{100}

	fmt.Println(hIndex(citations))
}
