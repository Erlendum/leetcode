package main

import "fmt"

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func main() {
	n := 2
	fmt.Println(climbStairs(n))
}
