package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTreeHelper(postorder []int, postorderIdx *int, inorderMap map[int]int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	nodeVal := postorder[*postorderIdx]
	node := new(TreeNode)
	node.Val = nodeVal
	*postorderIdx--

	inorderIdx := inorderMap[nodeVal]

	node.Right = buildTreeHelper(postorder, postorderIdx, inorderMap, inorderIdx+1, right)
	node.Left = buildTreeHelper(postorder, postorderIdx, inorderMap, left, inorderIdx-1)
	return node
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	m := make(map[int]int)
	for i, v := range inorder {
		m[v] = i
	}
	postorderIdx := len(postorder) - 1
	return buildTreeHelper(postorder, &postorderIdx, m, 0, len(inorder)-1)
}

func main() {
	postorder := []int{9, 15, 7, 20, 3}
	inorder := []int{9, 3, 15, 20, 7}

	fmt.Println(buildTree(inorder, postorder))
}
