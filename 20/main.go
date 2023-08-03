package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	data []byte
}

func Constructor() Stack {
	var s Stack
	s.data = make([]byte, 0)
	return s
}

func (this *Stack) Push(val byte) {
	this.data = append(this.data, val)
}

func (this *Stack) Pop() {
	this.data = this.data[:len(this.data)-1]
}

func (this *Stack) Top() byte {
	return this.data[len(this.data)-1]
}

func (this *Stack) IsEmpty() bool {
	return len(this.data) == 0
}

func isValid(s string) bool {
	opened := "([{"
	closed := ")]}"

	stack := Constructor()

	for i := range s {
		oi := strings.IndexByte(opened, s[i])
		ci := strings.IndexByte(closed, s[i])
		if oi >= 0 {
			stack.Push(closed[oi])
		} else if ci >= 0 && !stack.IsEmpty() && stack.Top() == s[i] {
			stack.Pop()
		} else {
			return false
		}
	}
	if !stack.IsEmpty() {
		return false
	}
	return true
}

func main() {
	s := "(]"
	fmt.Println(isValid(s))
}
