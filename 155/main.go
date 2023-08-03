package main

import (
	"fmt"
)

type MinStack struct {
	data []int
	mins []int
}

func Constructor() MinStack {
	var m MinStack
	m.data = make([]int, 0)
	m.mins = make([]int, 0)
	return m
}

func (this *MinStack) Push(val int) {
	if len(this.mins) == 0 || val < this.mins[len(this.mins)-1] {
		this.mins = append(this.mins, val)
	} else {
		this.mins = append(this.mins, this.mins[len(this.mins)-1])
	}
	this.data = append(this.data, val)
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.mins = this.mins[:len(this.mins)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}

func main() {
	m := Constructor()
	m.Push(1)
	m.Push(2)
	m.Push(3)
	fmt.Println(m.GetMin())
}
