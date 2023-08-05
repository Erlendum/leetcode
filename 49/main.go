package main

import (
	"fmt"
	"sort"
	"strings"
)

func createHash(m map[string]int) string {
	keys := make([]string, 0)
	for key, value := range m {
		keys = append(keys, key+string(rune(value)))
	}
	sort.Strings(keys)
	return strings.Join(keys, "")
}

func createMap(s string) map[string]int {
	m := make(map[string]int, 0)
	for i := range s {
		if _, ok := m[string(s[i])]; !ok {
			m[string(s[i])] = 1
		} else {
			m[string(s[i])]++
		}
	}
	return m
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string, 0)

	for i := range strs {
		sMap := createMap(strs[i])
		hash := createHash(sMap)
		if _, ok := m[createHash(sMap)]; !ok {
			m[hash] = make([]string, 0)
		}
		m[hash] = append(m[hash], strs[i])
	}
	res := make([][]string, len(m))
	i := 0
	for _, v := range m {
		res[i] = v
		i++
	}
	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
