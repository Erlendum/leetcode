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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	var bufList1, bufList2 *ListNode
	isList1 := true
	if list1.Val <= list2.Val {
		bufList1 = list1
		bufList2 = list2
	} else {
		bufList1 = list2
		bufList2 = list1
		isList1 = false
	}

	for bufList1 != nil && bufList2 != nil {
		tmp2 := bufList2.Next
		if bufList1.Next != nil && bufList1.Next.Val > bufList2.Val {
			tmp := bufList1.Next
			bufList1.Next = bufList2
			bufList2.Next = tmp
			bufList1 = bufList2
			bufList2 = tmp2
		} else if bufList1.Next == nil && bufList2 != nil {
			bufList1.Next = bufList2
			break
		} else {
			bufList1 = bufList1.Next
		}
	}
	if isList1 {
		return list1
	}
	return list2
}

func main() {
	l1 := toList([]int{2})
	l2 := toList([]int{1})
	Print(l1)
	fmt.Println()
	Print(l2)
	fmt.Println()
	newL := mergeTwoLists(l1, l2)
	Print(newL)
}
