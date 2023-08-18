package main

import (
	"container/list"
	"fmt"
	"strconv"
)

type customStack struct {
	stack *list.List
}

func (c *customStack) Push(value int) {
	c.stack.PushFront(value)
}

func (c *customStack) Pop() error {
	if c.stack.Len() > 0 {
		ele := c.stack.Front()
		c.stack.Remove(ele)
	}
	return fmt.Errorf("Pop Error: Stack is empty")
}

func (c *customStack) Front() (int, error) {
	if c.stack.Len() > 0 {
		if val, ok := c.stack.Front().Value.(int); ok {
			return val, nil
		}
		return 0, fmt.Errorf("Peep Error: Stack Datatype is incorrect")
	}
	return 0, fmt.Errorf("Peep Error: Stack is empty")
}

func (c *customStack) Size() int {
	return c.stack.Len()
}

func (c *customStack) Empty() bool {
	return c.stack.Len() == 0
}

func evalRPN(tokens []string) int {
	stack := customStack{new(list.List)}

	var n int
	for i := range tokens {
		v, err := strconv.Atoi(tokens[i])
		if err == nil {
			stack.Push(v)
		} else {
			r, _ := stack.Front()
			stack.Pop()
			l, _ := stack.Front()
			stack.Pop()
			switch tokens[i] {
			case "+":
				n = l + r
			case "-":
				n = l - r
			case "*":
				n = l * r
			case "/":
				n = l / r
			}
			stack.Push(n)
		}
	}
	res, _ := stack.Front()
	return res
}

func main() {
	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens))
}
