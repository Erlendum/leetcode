package main

import "fmt"

func hammingWeight(num uint32) int {
	var res int
	for num > 0 {
		if num%2 == 1 {
			res++
		}
		num /= 2
	}
	return res
}

func main() {
	var n uint32
	n = 11
	fmt.Println(hammingWeight(n))
}
