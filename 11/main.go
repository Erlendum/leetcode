package main

import "fmt"

func myMax(n1, n2 int) int {
	if n1 < n2 {
		return n2
	}
	return n1
}

func myMin(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxAr := 0
	for l < r {
		area := (r - l) * myMin(height[l], height[r])
		maxAr = myMax(area, maxAr)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxAr
}

func main() {
	fmt.Print(maxArea([]int{1, 1}))
}
