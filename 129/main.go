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

func sumNumbersHelper(root *TreeNode, prevRoot *TreeNode, num int, nums *[]int) {
	if root == nil {
		if prevRoot.Left == nil && prevRoot.Right == nil {
			*nums = append(*nums, num)
		}
		return
	}

	sumNumbersHelper(root.Left, root, num*10+root.Val, nums)
	sumNumbersHelper(root.Right, root, num*10+root.Val, nums)
}

func sumNumbers(root *TreeNode) int {
	nums := make([]int, 0)
	sumNumbersHelper(root, root, 0, &nums)
	s := 0
	for i := 0; i < len(nums); i += 2 {
		s += nums[i]
	}
	return s
}

func main() {
	tree := createTree([]string{"1", "0", "null"})
	Print(tree)
	fmt.Println()
	fmt.Println(sumNumbers(tree))
}
