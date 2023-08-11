package main

import "fmt"

type Pair struct {
	digit   int
	letters string
}

var pairs = []Pair{
	{2, "abc"},
	{3, "def"},
	{4, "ghi"},
	{5, "jkl"},
	{6, "mno"},
	{7, "pqrs"},
	{8, "tuv"},
	{9, "wxyz"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	digitsPointers := make([]int, len(digits))
	for i := range digitsPointers {
		digitsPointers[i] = 0
	}

	res := make([]string, 0)

	l := 1
	for i := range digits {
		l *= len(pairs[digits[i]-'0'-2].letters)
	}

	for i := 0; i < l; i++ {
		var s string
		for j := range digits {
			if digitsPointers[j] != -1 {
				s += string(pairs[digits[j]-'0'-2].letters[digitsPointers[j]])
			}
		}
		flag := false
		for j := len(digits) - 1; j >= 0; j-- {
			if j == len(digits)-1 || flag {
				digitsPointers[j]++
			}
			if digitsPointers[j] == len(pairs[digits[j]-'0'-2].letters) {
				digitsPointers[j] = 0
				flag = true
			} else {
				flag = false
			}
		}
		res = append(res, s)
	}
	return res
}

func main() {
	digits := "7"
	fmt.Println(letterCombinations(digits))
}
