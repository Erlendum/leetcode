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
	fmt.Printf("%d ", root.Val)
	Print(root.Left)
	Print(root.Right)
}

func kthSmallestHelper(root *TreeNode, k *int) int {
	if root == nil {
		return -1
	}
	if tmp := kthSmallestHelper(root.Left, k); tmp != -1 {
		return tmp
	}
	*k--
	if *k == 0 {
		return root.Val
	}
	return kthSmallestHelper(root.Right, k)
}

func kthSmallest(root *TreeNode, k int) int {
	return kthSmallestHelper(root, &k)
}

func main() {
	tree := createTree([]string{"3", "1", "4", "null", "2", "null", "null"})
	Print(tree)
	fmt.Println()
	k := 1
	fmt.Println(kthSmallest(tree, k))
}
