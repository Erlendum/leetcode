package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTreeHelper(preorder []int, preorderIdx *int, inorderMap map[int]int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	nodeVal := preorder[*preorderIdx]
	node := new(TreeNode)
	node.Val = nodeVal
	*preorderIdx++

	inorderIdx := inorderMap[nodeVal]

	node.Left = buildTreeHelper(preorder, preorderIdx, inorderMap, left, inorderIdx-1)
	node.Right = buildTreeHelper(preorder, preorderIdx, inorderMap, inorderIdx+1, right)
	return node
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	m := make(map[int]int)
	for i, v := range inorder {
		m[v] = i
	}
	var preorderIdx int
	return buildTreeHelper(preorder, &preorderIdx, m, 0, len(inorder)-1)
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	fmt.Println(buildTree(preorder, inorder))
}
