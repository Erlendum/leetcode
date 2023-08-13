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

type BSTIterator struct {
	values []int
	next   int
}

func Constructor(root *TreeNode) BSTIterator {
	iterator := new(BSTIterator)
	iterator.values = make([]int, 0)
	inOrderAppend(root, &iterator.values)
	iterator.next = 0
	return *iterator
}

func inOrderAppend(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}

	inOrderAppend(root.Left, values)
	*values = append(*values, root.Val)
	inOrderAppend(root.Right, values)
}

func (this *BSTIterator) Next() int {
	prevNext := this.next
	this.next++
	return this.values[prevNext]
}

func (this *BSTIterator) HasNext() bool {
	return this.next < len(this.values)
}
