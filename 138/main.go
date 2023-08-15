package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func Println(head *Node) {
	for ; head != nil; head = head.Next {
		fmt.Printf("%d ", head.Val)
	}
	fmt.Println()
}

func copyRandomListHelper(head *Node, m map[*Node]*Node) *Node {
	if head == nil {
		return nil
	}
	if _, ok := m[head]; ok {
		return m[head]
	}

	node := &Node{Val: head.Val}
	m[head] = node
	node.Random = copyRandomListHelper(head.Random, m)
	node.Next = copyRandomListHelper(head.Next, m)
	return node
}

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	return copyRandomListHelper(head, m)
}

func main() {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n1.Next = n2
	n1.Random = nil
	n2.Next = nil
	n2.Random = n2
	Println(n1)
	newL := copyRandomList(n1)
	Println(newL)
}
