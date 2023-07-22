package main

import "fmt"

func find(slice []string, s string, start int, stop int) int {
	if start < 0 || stop > len(s) {
		return -1
	}
	for i := range slice {
		if slice[i] == s[start:stop] {
			return i
		}
	}
	return -1
}
func romanToInt(s string) int {
	s1 := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	s2 := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	s3 := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	s4 := []string{"", "M", "MM", "MMM", "MMMM"}

	res := 0
	index := 0

	for index < len(s) {
		if s[index] == 'I' || s[index] == 'V' {
			if i := find(s1, s, index, index+4); i > 0 {
				res += i
				index += 4
			} else if i := find(s1, s, index, index+3); i > 0 {
				res += i
				index += 3
			} else if i := find(s1, s, index, index+2); i > 0 {
				res += i
				index += 2
			} else if i := find(s1, s, index, index+1); i > 0 {
				res += i
				index += 1
			}
		} else if s[index] == 'X' || s[index] == 'L' {
			if i := find(s2, s, index, index+4); i > 0 {
				res += i * 10
				index += 4
			} else if i := find(s2, s, index, index+3); i > 0 {
				res += i * 10
				index += 3
			} else if i := find(s2, s, index, index+2); i > 0 {
				res += i * 10
				index += 2
			} else if i := find(s2, s, index, index+1); i > 0 {
				res += i * 10
				index += 1
			}
		} else if s[index] == 'C' || s[index] == 'D' {
			if i := find(s3, s, index, index+4); i > 0 {
				res += i * 100
				index += 4
			} else if i := find(s3, s, index, index+3); i > 0 {
				res += i * 100
				index += 3
			} else if i := find(s3, s, index, index+2); i > 0 {
				res += i * 100
				index += 2
			} else if i := find(s3, s, index, index+1); i > 0 {
				res += i * 100
				index += 1
			}
		} else if s[index] == 'M' {
			if i := find(s4, s, index, index+4); i > 0 {
				res += i * 1000
				index += 4
			} else if i := find(s4, s, index, index+3); i > 0 {
				res += i * 1000
				index += 3
			} else if i := find(s4, s, index, index+2); i > 0 {
				res += i * 1000
				index += 2
			} else if i := find(s4, s, index, index+1); i > 0 {
				res += i * 1000
				index += 1
			}
		}
	}
	return res
}

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}
