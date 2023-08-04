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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := Length(head)
	n = length - n

	if n == 0 {
		return head.Next
	}
	bufHead := head
	if n == length-1 {
		for bufHead.Next.Next != nil {
			bufHead = bufHead.Next
		}
		bufHead.Next = nil
		return head
	}

	for i := 0; i < n-1; i++ {
		bufHead = bufHead.Next
	}
	bufHead.Next = bufHead.Next.Next
	return head
}

func main() {
	l := toList([]int{1})
	Print(l)
	fmt.Println()
	newL := removeNthFromEnd(l, 1)
	Print(newL)
}
