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

func rotateRight(head *ListNode, k int) *ListNode {
	length := Length(head)
	if length == 0 || length == 1 || k%length == 0 {
		return head
	}
	initHead := head
	stop := head

	for i := 0; i < length-k%length-1; i++ {
		stop = stop.Next
	}

	start := stop.Next

	for head.Next != nil {
		head = head.Next
	}
	head.Next = initHead
	stop.Next = nil
	return start
}

func main() {
	l := toList([]int{1, 2, 0})
	newL := rotateRight(l, 4)
	Print(newL)
}
