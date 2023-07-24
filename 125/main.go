package main

import (
	"fmt"
	"strings"
)

const marks = ",;.:!?\"<>~`@#№$%^&*()_+=-{}'[] "

func strip(s string) string {
	var res string
	for i := range s {
		if !strings.ContainsAny(marks, s[i:i+1]) {
			res += strings.ToLower(string(s[i]))
		}
	}
	return res
}

func isPalindrome(s string) bool {
	newS := strip(s)
	var p1, p2 int
	p2 = len(newS) - 1

	for p1 >= 0 && p2 >= 0 && p1 != p2 {
		if newS[p1] != newS[p2] {
			return false
		}
		p1++
		p2--
	}
	return true
}

func main() {
	s := "Marge, let's \"[went].\" I await {news} telegram."
	fmt.Println(isPalindrome(s))
}
