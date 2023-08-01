package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	m1 := make(map[uint8]int, 0)
	m2 := make(map[uint8]int, 0)

	for i := range ransomNote {
		if _, ok := m1[ransomNote[i]]; !ok {
			m1[ransomNote[i]] = 1
		} else {
			m1[ransomNote[i]]++
		}
	}

	for i := range magazine {
		if _, ok := m2[magazine[i]]; !ok {
			m2[magazine[i]] = 1
		} else {
			m2[magazine[i]]++
		}
	}

	for k := range m1 {
		if m1[k] > m2[k] {
			return false
		}
	}
	return true
}

func main() {
	ransomNote := "aa"
	magazine := "ab"
	fmt.Println(canConstruct(ransomNote, magazine))
}
