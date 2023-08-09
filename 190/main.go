package main

import (
	"fmt"
	"strconv"
)

func binReversed(num uint32) string {
	var res string
	for num > 0 {
		res = res + strconv.Itoa(int(num%2))
		num /= 2
	}
	for len(res) < 32 {
		res += "0"
	}
	return res
}

func myPow(n, i uint32) uint32 {
	var res uint32
	res = 1
	var k uint32
	for ; k < i; k++ {
		res *= n
	}
	return res
}
func toUint(bin string) uint32 {
	var res uint32

	for i := len(bin) - 1; i >= 0; i-- {
		if bin[i] == '1' {
			res += myPow(2, uint32(len(bin))-uint32(i)-1)
		}
	}
	return res
}
func reverseBits(num uint32) uint32 {
	return toUint(binReversed(num))
}

func main() {
	var n uint32
	n = 4294967293
	fmt.Println(reverseBits(n))
}
