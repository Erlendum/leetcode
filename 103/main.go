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

func FillMapWithDefaultValues(root *TreeNode, m map[*TreeNode]int) {
	if root == nil {
		return
	}
	m[root] = -1
	FillMapWithDefaultValues(root.Left, m)
	FillMapWithDefaultValues(root.Right, m)
}

func maxKey(m map[int][]int) int {
	myMax := 0
	for k := range m {
		if k > myMax {
			myMax = k
		}
	}
	return myMax
}

func reverse(slice []int) []int {
	res := make([]int, len(slice))
	j := 0
	for i := len(slice) - 1; i >= 0; i-- {
		res[j] = slice[i]
		j++
	}
	return res
}

func zigzagLevelOrderHelper(root *TreeNode, values *[][]int) {
	visited := make(map[*TreeNode]int)
	FillMapWithDefaultValues(root, visited)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	visited[root] = 1
	levelMap := make(map[int][]int)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if _, ok := levelMap[visited[node]]; !ok {
			levelMap[visited[node]] = make([]int, 0)
		}
		levelMap[visited[node]] = append(levelMap[visited[node]], node.Val)

		if node.Left != nil && visited[node.Left] == -1 {
			queue = append(queue, node.Left)
			visited[node.Left] = visited[node] + 1
		}
		if node.Right != nil && visited[node.Right] == -1 {
			queue = append(queue, node.Right)
			visited[node.Right] = visited[node] + 1
		}
	}

	mKey := maxKey(levelMap)
	valuesCopy := make([][]int, mKey+1)
	for k := range levelMap {
		if k%2 == 0 {
			valuesCopy[k] = reverse(levelMap[k])
		} else {
			valuesCopy[k] = levelMap[k]
		}
	}
	valuesCopy = valuesCopy[1:]
	*values = valuesCopy
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	zigzagLevelOrderHelper(root, &res)
	return res
}

func main() {
	tree := createTree([]string{"3", "9", "20", "null", "null", "15", "7"})
	Print(tree)
	fmt.Println()
	fmt.Println(zigzagLevelOrder(tree))
}
