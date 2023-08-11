package main

import (
	"fmt"
)

func trailingZeroes(n int) int {
	var c int

	for i := 5; i <= n; i *= 5 {
		c += n / i
	}
	return c
}

func main() {
	n := 7890

	fmt.Println(trailingZeroes(n))
}
