package main

import "fmt"

func strStr(haystack string, needle string) int {
	var p1, p1Tmp, p2 int

	for p1 < len(haystack) {
		if p2 == len(needle) {
			return p1
		}
		if p1Tmp >= len(haystack) {
			return -1
		}
		if haystack[p1Tmp] != needle[p2] {
			p2 = 0
			p1++
			p1Tmp = p1
		} else {
			p1Tmp++
			p2++
		}
	}
	return -1
}

func main() {
	haystack := "aaa"
	needle := "aaaa"
	fmt.Println(strStr(haystack, needle))
}
