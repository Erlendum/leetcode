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

func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]struct{}, 0)

	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return false
}

func main() {
	list := toList([]int{3, 2, 0, 4})
	list.Next.Next.Next.Next = list.Next
	fmt.Println(hasCycle(list))
}
