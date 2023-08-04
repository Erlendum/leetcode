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

func deleteDuplicates(head *ListNode) *ListNode {
	bufHead := head
	if bufHead == nil || bufHead.Next == nil {
		return head
	}
	flag := bufHead.Val == bufHead.Next.Val

	for bufHead.Next != nil {
		tmp := bufHead.Next.Next
		if tmp == nil || bufHead.Next.Val != tmp.Val {
			bufHead = bufHead.Next
			continue
		}
		for tmp != nil && bufHead.Next.Val == tmp.Val {
			tmp = tmp.Next
		}
		bufHead.Next = tmp
	}
	if flag {
		if head.Next != nil && head.Val == head.Next.Val {
			return head.Next.Next
		}
		return head.Next
	}
	return head
}

func main() {
	l := toList([]int{1, 1, 1})
	Print(l)
	fmt.Println()
	newL := deleteDuplicates(l)
	Print(newL)
}
