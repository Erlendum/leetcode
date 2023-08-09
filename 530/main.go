package main

import (
	"fmt"
	"math"
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

func myAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
func getMinimumDifferenceHelper(root *TreeNode, prevVal *int, min *int) {
	if root == nil {
		return
	}
	getMinimumDifferenceHelper(root.Left, prevVal, min)
	if *prevVal != -1 && myAbs(root.Val-*prevVal) < *min {
		*min = myAbs(root.Val - *prevVal)
	}

	*prevVal = root.Val
	getMinimumDifferenceHelper(root.Right, prevVal, min)
}

func getMinimumDifference(root *TreeNode) int {
	min := math.MaxInt
	prevVal := -1
	getMinimumDifferenceHelper(root, &prevVal, &min)
	return min
}

func main() {
	tree := createTree([]string{"236", "104", "701", "null", "227", "null", "911"})
	Print(tree)
	fmt.Println()
	fmt.Println(getMinimumDifference(tree))
}
