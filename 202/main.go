package main

import "fmt"

func sumSquaresOfDigits(n int) int {
	s := 0
	for n > 0 {
		s += (n % 10) * (n % 10)
		n /= 10
	}
	return s
}

func isHappy(n int) bool {
	if n <= 0 {
		return false
	}
	iterations := 50
	m := make(map[int]struct{}, 0)
	for i := 0; i < iterations && n != 1; i++ {
		if _, ok := m[n]; !ok {
			m[n] = struct{}{}
		} else {
			return false
		}
		n = sumSquaresOfDigits(n)
	}
	if n == 1 {
		return true
	}
	return false
}

func main() {
	n := 2
	fmt.Println(isHappy(n))
}
