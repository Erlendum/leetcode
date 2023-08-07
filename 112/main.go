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
		for j := i; j < len(data) && j <= 2*i; j++ {
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

func hasPathSumHelper(root *TreeNode, prevRoot *TreeNode, targetSum int, curSum int) bool {
	if root == nil {
		if prevRoot.Left != nil || prevRoot.Right != nil {
			return false
		}
		return curSum == targetSum
	}

	if hasPathSumHelper(root.Left, root, targetSum, curSum+root.Val) {
		return true
	}

	if hasPathSumHelper(root.Right, root, targetSum, curSum+root.Val) {
		return true
	}

	return false
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return hasPathSumHelper(root, root, targetSum, 0)
}

func main() {
	tree := new(TreeNode)
	tree.Val = 1
	tree.Left = new(TreeNode)
	tree.Left.Val = 2
	tree.Left.Left = new(TreeNode)
	tree.Left.Left.Val = 3
	tree.Left.Left.Left = new(TreeNode)
	tree.Left.Left.Left.Val = 4
	tree.Left.Left.Left.Left = new(TreeNode)
	tree.Left.Left.Left.Left.Val = 5
	target := 5
	Print(tree)
	fmt.Println()
	fmt.Println(hasPathSum(tree, target))
}
