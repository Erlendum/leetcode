package main

import "fmt"

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	l := 0
	for ; i >= 0 && s[i] == ' '; i-- {

	}
	for ; i >= 0 && s[i] != ' '; i-- {
		l++
	}
	return l
}

func main() {
	s := "a"
	fmt.Println(lengthOfLastWord(s))
}
