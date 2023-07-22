package main

import (
	"fmt"
	"os"
)

func isMatch(s string, p string) bool {
	sIndex := len(s) - 1
	pIndex := len(p) - 1

	for sIndex >= 0 && pIndex >= 0 {
		if p[pIndex] != '.' && p[pIndex] != '*' {
			if s[sIndex] != p[pIndex] {
				return false
			} else {
				sIndex--
				pIndex--
			}
		} else if p[pIndex] == '.' {
			sIndex--
			pIndex--
		} else if p[pIndex] == '*' {
			if s[sIndex] == p[pIndex-1] || p[pIndex-1] == '.' {
				if isMatch(s[:sIndex+1], p[:pIndex-1]) {
					return true
				}
				sIndex--
			}
			if sIndex < 0 || (s[sIndex] != p[pIndex-1] && p[pIndex-1] != '.') {
				pIndex -= 2
			}
		}
	}

	for pIndex >= 0 {
		if p[pIndex] != '*' {
			return false
		}
		pIndex -= 2
	}
	if sIndex < 0 && pIndex < 0 {
		return true
	}
	return false
}

func main() {
	var s, p string

	_, err := fmt.Scan(&s)
	if err != nil {
		os.Exit(1)
	}

	_, err = fmt.Scan(&p)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(isMatch(s, p))
}
