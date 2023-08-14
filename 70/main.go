package main

import "fmt"

func climbStairsHelper(n int, memo map[int]int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if _, ok := memo[n]; !ok {
		memo[n] = climbStairsHelper(n-1, memo) + climbStairsHelper(n-2, memo)
	}
	return memo[n]
}

func climbStairs(n int) int {
	memo := make(map[int]int)
	return climbStairsHelper(n, memo)
}

func main() {
	n := 2
	fmt.Println(climbStairs(n))
}
