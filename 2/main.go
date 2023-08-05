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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resList := Constructor(0)
	resListBuf := resList
	resListPrevLast := resList
	var rem int
	for l1 != nil || l2 != nil {
		var l1Val, l2Val int
		if l1 != nil {
			l1Val = l1.Val
		}
		if l2 != nil {
			l2Val = l2.Val
		}
		res := (l1Val + l2Val + rem) % 10
		rem = (l1Val + l2Val + rem) / 10

		resListBuf.Val = res
		resListBuf.Next = Constructor(0)
		resListPrevLast = resListBuf
		resListBuf = resListBuf.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if rem != 0 {
		resListBuf.Val = rem
	}
	if resListPrevLast.Next.Val == 0 {
		resListPrevLast.Next = nil
	}

	return resList
}

func main() {
	l1 := toList([]int{0})
	l2 := toList([]int{0})
	Print(l1)
	fmt.Println()
	Print(l2)
	fmt.Println()
	res := addTwoNumbers(l1, l2)
	Print(res)
}
