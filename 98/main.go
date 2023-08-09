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

type Pair struct {
	val     int
	isRight bool
}

func myCmp(val int, values []Pair) bool {
	for i := range values {
		if values[i].isRight && val <= values[i].val {
			return false
		}
		if !values[i].isRight && val >= values[i].val {
			return false
		}
	}
	return true
}
func isValidBSTHelper(root *TreeNode, path []Pair) bool {
	if root == nil {
		return true
	}
	if (root.Left != nil && root.Left.Val >= root.Val) || (root.Right != nil && root.Right.Val <= root.Val) {
		return false
	}

	if root.Left != nil && !myCmp(root.Left.Val, path) {
		return false
	}
	if root.Right != nil && !myCmp(root.Right.Val, path) {
		return false
	}
	path = append(path, Pair{root.Val, false})
	if !isValidBSTHelper(root.Left, path) {
		return false
	}
	path[len(path)-1].isRight = true
	if !isValidBSTHelper(root.Right, path) {
		return false
	}
	return true
}
func isValidBST(root *TreeNode) bool {
	path := make([]Pair, 0)
	return isValidBSTHelper(root, path)
}

func main() {
	tree := new(TreeNode)
	tree.Val = 3
	tree.Left = new(TreeNode)
	tree.Left.Val = 1
	tree.Right = new(TreeNode)
	tree.Right.Val = 5
	tree.Left.Left = new(TreeNode)
	tree.Left.Left.Val = 0
	tree.Left.Right = new(TreeNode)
	tree.Left.Right.Val = 2
	tree.Left.Right.Right = new(TreeNode)
	tree.Left.Right.Right.Val = 3
	Print(tree)
	fmt.Println()
	fmt.Println(isValidBST(tree))

}
