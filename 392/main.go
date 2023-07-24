package main

import "fmt"

func isSubsequence(s string, t string) bool {
	var p1, p2 int

	for p1 < len(t) && p2 < len(s) {
		if t[p1] == s[p2] {
			p2++
		}
		p1++
	}

	return p2 == len(s)
}

func main() {
	s := ""
	t := "ahbgdc"

	fmt.Println(isSubsequence(s, t))
}
