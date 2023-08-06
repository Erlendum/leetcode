package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createNode(val int) *TreeNode {
	node := new(TreeNode)
	node.Val = val
	node.Left = nil
	node.Right = nil
	return node
}

func createTree(data []string) *TreeNode {
	rootVal, err := strconv.Atoi(data[0])
	if err != nil {
		return nil
	}
	root := createNode(rootVal)

	curNodes := make([]*TreeNode, 0)
	curNodes = append(curNodes, root)
	for i := 1; i < len(data); i += 2 * i {
		for j := i; j <= 2*i; j++ {
			var node *TreeNode
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

func Print(root *TreeNode) {
	if root == nil {
		fmt.Print("null ")
		return
	}
	Print(root.Left)
	Print(root.Right)
	fmt.Printf("%d ", root.Val)
}

func maxDepthHelper(root *TreeNode, m int, myMax *int) {
	if root == nil {
		if m > *myMax {
			*myMax = m
		}
		return
	}
	maxDepthHelper(root.Left, m+1, myMax)
	maxDepthHelper(root.Right, m+1, myMax)
}

func maxDepth(root *TreeNode) int {
	myMax := new(int)
	maxDepthHelper(root, 0, myMax)
	return *myMax
}

func main() {
	tree := createTree([]string{"3", "9", "20", "null", "null", "15", "7"})
	Print(tree)
	fmt.Println()
	fmt.Println(maxDepth(tree))
}
