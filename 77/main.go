package main

import "fmt"

func fact(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func countCombinations(n, k int) int {
	return fact(n) / (fact(k) * fact(n-k))
}

func combine(n int, k int) [][]int {
	integers := make([]int, n)
	for i := 0; i < n; i++ {
		integers[i] = i + 1
	}

	integersPointers := make([]int, k)
	for i := 0; i < len(integersPointers); i++ {
		integersPointers[i] = i
	}

	res := make([][]int, 0)

	l := countCombinations(n, k)

	for i := 0; i < l; i++ {
		el := make([]int, 0)
		for j := range integersPointers {
			el = append(el, integers[integersPointers[j]])
		}

		j := len(integersPointers) - 1
		integersPointers[j]++
		if j != 0 && integersPointers[j] >= len(integers) {
			tmpJ := j
			for ; tmpJ > 0 && integersPointers[j] >= len(integers); tmpJ-- {
				integersPointers[tmpJ-1]++
				for k := tmpJ; k < len(integersPointers); k++ {
					integersPointers[k] = integersPointers[k-1] + 1
				}
			}
		}

		res = append(res, el)
	}
	return res
}

func main() {
	n := 5
	k := 4
	fmt.Println(combine(n, k))
}
