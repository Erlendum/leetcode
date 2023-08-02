package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := make(map[uint8]uint8, 0)
	m2 := make(map[uint8]uint8, 0)

	for i := range s {
		_, ok1 := m1[s[i]]
		_, ok2 := m2[t[i]]
		if !ok1 && !ok2 {
			m1[s[i]] = t[i]
			m2[t[i]] = s[i]
		} else {
			if m1[s[i]] != t[i] || m2[t[i]] != s[i] {
				return false
			}
		}
	}
	return true
}

func main() {
	s := "ega"
	t := "add"

	fmt.Println(isIsomorphic(s, t))
}
