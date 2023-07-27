package main

import "fmt"

func mySqrt(x int) int {
	if x == 1 {
		return x
	}
	i := 0
	for ; i <= x/2; i++ {
		if i*i == x {
			return i
		} else if i*i > x {
			return i - 1
		}
	}
	return i - 1
}

func main() {
	x := 2

	fmt.Println(mySqrt(x))
}
