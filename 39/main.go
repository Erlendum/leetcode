package main

import (
	"fmt"
	"sort"
)

func removeDuplicateValues(intSlice []int) []int {
	keys := make(map[int]bool)
	var list []int

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func combinationSumHelper(candidates []int, target int, cur *[]int, res *[][]int, i int) {
	if target == 0 {
		tmp := make([]int, len(*cur))
		copy(tmp, *cur)
		*res = append(*res, tmp)
		return
	}

	for ; i < len(candidates) && target-candidates[i] >= 0; i++ {
		*cur = append(*cur, candidates[i])
		combinationSumHelper(candidates, target-candidates[i], cur, res, i)
		*cur = (*cur)[:len(*cur)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	removeDuplicateValues(candidates)

	sort.Ints(candidates)

	res := make([][]int, 0)
	cur := make([]int, 0)

	combinationSumHelper(candidates, target, &cur, &res, 0)
	return res
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7

	fmt.Println(combinationSum(candidates, target))
}
