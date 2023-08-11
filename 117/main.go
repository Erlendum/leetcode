package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func createNode(val int) *Node {
	node := new(Node)
	node.Val = val
	node.Left = nil
	node.Right = nil
	return node
}

func createTree(data []string) *Node {
	rootVal, err := strconv.Atoi(data[0])
	if err != nil {
		return nil
	}
	root := createNode(rootVal)

	curNodes := make([]*Node, 0)
	curNodes = append(curNodes, root)
	for i := 1; i < len(data); i += 2 * i {
		for j := i; j <= 2*i; j++ {
			var node *Node
			if data[j] != "null" {
				nodeVal, err := strconv.Atoi(data[j])
				if err != nil {
					return nil
				}
				node = createNode(nodeVal)
			}
			if j%2 == 1 {
				if curNodes[0] != nil {
					curNodes[0].Left = node
				}
			} else {
				if curNodes[0] != nil {
					curNodes[0].Right = node
				}
				curNodes = curNodes[1:]
			}
			curNodes = append(curNodes, node)

		}
	}
	return root
}

func Print(root *Node) {
	if root == nil {
		fmt.Print("null ")
		return
	}
	fmt.Printf("%d ", root.Val)
	Print(root.Left)
	Print(root.Right)
}

func FillMapWithDefaultValues(root *Node, m map[*Node]int) {
	if root == nil {
		return
	}
	m[root] = -1
	FillMapWithDefaultValues(root.Left, m)
	FillMapWithDefaultValues(root.Right, m)
}

func connectHelper(root *Node) {
	visited := make(map[*Node]int)
	FillMapWithDefaultValues(root, visited)
	queue := make([]*Node, 0)
	queue = append(queue, root)
	visited[root] = 1
	prevRoot := root
	for len(queue) != 0 {
		node := queue[0]
		if len(queue) > 1 {
			node.Next = queue[1]
		}
		if visited[prevRoot] != visited[node] {
			prevRoot.Next = nil
		}
		prevRoot = node
		queue = queue[1:]

		if node.Left != nil && visited[node.Left] == -1 {
			queue = append(queue, node.Left)
			visited[node.Left] = visited[node] + 1
		}
		if node.Right != nil && visited[node.Right] == -1 {
			queue = append(queue, node.Right)
			visited[node.Right] = visited[node] + 1
		}
	}
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectHelper(root)
	return root
}

func main() {
	tree := createTree([]string{"1", "2", "3", "4", "5", "null", "7"})
	Print(tree)
	fmt.Println()
	fmt.Println(connect(tree))
}
