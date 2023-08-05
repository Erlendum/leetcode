package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	var res string
	ia := len(a) - 1
	ib := len(b) - 1
	var rem int
	for ia >= 0 || ib >= 0 {
		var int1, int2 int
		if ia >= 0 {
			int1, _ = strconv.Atoi(string(a[ia]))
		}
		if ib >= 0 {
			int2, _ = strconv.Atoi(string(b[ib]))
		}
		val := (int1 + int2 + rem) % 2
		rem = (int1 + int2 + rem) / 2
		res = strconv.Itoa(val) + res
		ia--
		ib--
	}
	if rem != 0 {
		res = strconv.Itoa(rem) + res
	}
	return res
}

func main() {
	a := "1010"
	b := "1011"
	fmt.Println(addBinary(a, b))
}
