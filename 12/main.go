package main

import "fmt"

func intToRoman(num int) string {
	s1 := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	s2 := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	s3 := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	s4 := []string{"", "M", "MM", "MMM", "MMMM"}

	p1 := s4[num/1000]
	p2 := s3[num/100%10]
	p3 := s2[num/10%10]
	p4 := s1[num%10]

	return p1 + p2 + p3 + p4

}

func main() {
	num := 123
	fmt.Println(intToRoman(num))
}
