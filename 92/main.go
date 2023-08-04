package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Constructor(val int) *ListNode {
	var node ListNode
	node.Val = val
	node.Next = nil
	return &node
}

func toList(slice []int) *ListNode {
	var head *ListNode
	var tmp *ListNode
	for i := len(slice) - 1; i >= 0; i-- {
		head = Constructor(slice[i])
		head.Next = tmp
		tmp = head
	}
	return head
}

func Length(head *ListNode) int {
	l := 0
	for head != nil {
		l++
		head = head.Next
	}
	return l
}

func Print(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	bufHead := head
	startI := 0
	var prevLeft *ListNode
	for i := 0; i < left-1; i++ {
		if i == left-1-1 {
			prevLeft = head
		}
		head = head.Next
		startI++
	}

	slice := make([]*ListNode, 0)

	for i := 0; i < right-startI; i++ {
		slice = append(slice, head)
		head = head.Next
	}

	slice[0].Next = head
	if prevLeft != nil {
		prevLeft.Next = slice[len(slice)-1]
	} else {
		bufHead = slice[len(slice)-1]
	}
	for i := len(slice) - 1; i > 0; i-- {
		slice[i].Next = slice[i-1]
	}
	return bufHead
}

func main() {
	l := toList([]int{1, 2, 3, 4, 5})
	Print(l)
	fmt.Println()
	newL := reverseBetween(l, 1, 5)
	Print(newL)
}
