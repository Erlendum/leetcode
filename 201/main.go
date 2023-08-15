package main

import (
	"fmt"
)

func rangeBitwiseAnd(left int, right int) int {
	res := right

	for right > left {
		res = right & (right - 1)
		right = res
	}
	return res
}

func main() {
	left := 5
	right := 7
	fmt.Println(rangeBitwiseAnd(left, right))
}
