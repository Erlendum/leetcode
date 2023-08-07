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

func flattenHelper(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}
	*nodes = append(*nodes, root)
	flattenHelper(root.Left, nodes)
	flattenHelper(root.Right, nodes)
}

func flatten(root *TreeNode) {
	nodes := make([]*TreeNode, 0)
	flattenHelper(root, &nodes)
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Left = nil
		nodes[i].Right = nodes[i+1]
	}
}

func main() {
	tree := createTree([]string{"1", "2", "5", "3", "4", "null", "6"})
	Print(tree)
	fmt.Println()
	flatten(tree)
	Print(tree)
}
