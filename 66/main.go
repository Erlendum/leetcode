package main

import "fmt"

func plusOne(digits []int) []int {
	r := 1
	res := make([]int, 0)
	i := len(digits) - 1
	for ; i >= 0; i-- {
		res = append([]int{(digits[i] + r) % 10}, res...)
		r = (digits[i] + r) / 10
		if r == 0 {
			break
		}
	}
	if i != -1 {
		res = append(digits[:i], res...)
	}
	if r != 0 {
		res = append([]int{r}, res...)
	}
	return res
}

func main() {
	digits := []int{8, 9, 9, 9}

	fmt.Println(plusOne(digits))
}
