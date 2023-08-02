package main

import "fmt"

func split(s string) []string {
	res := make([]string, 0)

	p1 := 0
	p2 := 0
	for p1 < len(s) {
		if p2 < len(s) && s[p2] != ' ' {
			p2++
		} else {
			res = append(res, s[p1:p2])
			p2++
			p1 = p2
		}
	}
	return res
}

func wordPattern(pattern string, s string) bool {
	words := split(s)
	if len(words) != len(pattern) {
		return false
	}
	m := make(map[uint8]string, 0)
	mExtra := make(map[string]uint8, 0)

	wordi := 0
	for i := range pattern {
		_, ok1 := m[pattern[i]]
		_, ok2 := mExtra[words[wordi]]
		if !ok1 && !ok2 {
			m[pattern[i]] = words[wordi]
			mExtra[words[wordi]] = pattern[i]
		} else {
			if m[pattern[i]] != words[wordi] {
				return false
			}
		}
		wordi++
	}
	return true
}

func main() {
	pattern := "aaa"
	s := "aa aa aa aa"
	fmt.Println(wordPattern(pattern, s))
}
