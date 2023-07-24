package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	p1 := 0
	p2 := len(s) - 1

	for p1 < p2 {
		if s[p1] != s[p2] {
			return false
		}
		p1++
		p2--
	}
	return true
}

func main() {
	x := 11
	fmt.Println(isPalindrome(x))
}
