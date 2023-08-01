package main

import "fmt"

func isAnagram(s string, t string) bool {
	m1 := make(map[uint8]int, 0)
	m2 := make(map[uint8]int, 0)

	for i := range s {
		if _, ok := m1[s[i]]; !ok {
			m1[s[i]] = 1
		} else {
			m1[s[i]]++
		}
	}

	for i := range t {
		if _, ok := m2[t[i]]; !ok {
			m2[t[i]] = 1
		} else {
			m2[t[i]]++
		}
	}

	if len(m1) != len(m2) {
		return false
	}

	for k := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}
	return true
}

func main() {
	s := "anagram"
	t := "nagaram"

	fmt.Println(isAnagram(s, t))
}
