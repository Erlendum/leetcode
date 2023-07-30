package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	i := 1
	curI := 1
	curLen := 1
	maxLen := 0
	myMap := make(map[uint8]struct{}, 0)
	myMap[s[curI-1]] = struct{}{}
	for curI < len(s) && i < len(s) {
		if _, ok := myMap[s[curI]]; !ok {
			myMap[s[curI]] = struct{}{}
			curI++
			curLen++
		} else {
			i++
			curI = i
			if curLen > maxLen {
				maxLen = curLen
			}
			myMap = make(map[uint8]struct{}, 0)
			myMap[s[curI-1]] = struct{}{}
			curLen = 1
		}
	}
	if curLen > maxLen {
		maxLen = curLen
	}
	return maxLen
}

func main() {
	s := ""

	fmt.Println(lengthOfLongestSubstring(s))
}
