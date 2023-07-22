package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	index := 0

	if len(strs) == 1 {
		return strs[0]
	}
	for {
		for i := 1; i < len(strs); i++ {
			if index >= len(strs[i]) && index >= len(strs[i-1]) {
				return strs[0][:index]
			} else if index < len(strs[i]) && index < len(strs[i-1]) && strs[i][index] != strs[i-1][index] {
				return strs[0][:index]
			} else if (index < len(strs[i]) && index >= len(strs[i-1])) || (index >= len(strs[i]) && index < len(strs[i-1])) {
				return strs[0][:index]
			}
		}
		index++
	}
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
