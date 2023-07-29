package main

import (
	"fmt"
	"math"
)

const EPS = 1e-15

func isEqual(a, b float64) bool {
	return math.Abs(a-b) <= EPS
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
func myPow(x float64, n int) float64 {
	if n == 0 || isEqual(x, 1.0) {
		return 1.0
	}
	init := x
	res := 1.0
	for i := Abs(n); i > 0; {
		if i%2 == 0 {
			init = init * init
			i /= 2
		} else {
			res *= init
			i--
		}
	}
	if n < 0 {
		return 1.0 / res
	}
	return res
}

func main() {
	x := 2.0
	n := 10
	fmt.Println(myPow(x, n))
}
