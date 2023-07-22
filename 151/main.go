package main

import "fmt"

func split(s string) []string {
	res := make([]string, 0)

	p1, p2 := 0, 0
	for p1 < len(s) || p2 < len(s) {
		if p1 < len(s) && p2 < len(s) && s[p1] == ' ' && s[p2] == ' ' {
			p1++
			p2++
		} else if p1 < len(s) && p2 < len(s) && s[p1] != ' ' && s[p2] != ' ' {
			p2++
		} else if p2 >= len(s) || s[p2] == ' ' {
			res = append(res, s[p1:p2])
			p1 = p2
		}
	}

	return res
}

func reverseWords(s string) string {
	words := split(s)
	res := ""
	for i := len(words) - 1; i >= 0; i-- {
		res += words[i]
		if i != 0 {
			res += " "
		}
	}
	return res
}

func main() {
	s := "a good   example"
	fmt.Println(reverseWords(s))
}
