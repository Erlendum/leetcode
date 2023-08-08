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

type AVGData struct {
	sum   int
	count int
}

const CountOfLevels = 500

func averageOfLevelsHelper(root *TreeNode, values *[]float64) {
	visited := make(map[*TreeNode]int)
	FillMapWithDefaultValues(root, visited)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	visited[root] = 1

	avgData := make([]AVGData, CountOfLevels)
	for len(queue) != 0 {
		node := queue[0]
		avgData[visited[node]] = AVGData{sum: avgData[visited[node]].sum + node.Val, count: avgData[visited[node]].count + 1}
		queue = queue[1:]

		if node.Left != nil && visited[node.Left] == -1 {
			queue = append(queue, node.Left)
			visited[node.Left] = visited[node] + 1
		}
		if node.Right != nil && visited[node.Right] == -1 {
			queue = append(queue, node.Right)
			visited[node.Right] = visited[node] + 1
		}
	}
	for _, c := range avgData {
		if c.count != 0 {
			*values = append(*values, float64(c.sum)/float64(c.count))
		}
	}
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	res := make([]float64, 0)
	averageOfLevelsHelper(root, &res)
	return res
}

func main() {
	tree := createTree([]string{"3", "9", "20", "null", "null", "15", "7"})
	Print(tree)
	fmt.Println()
	fmt.Println(averageOfLevels(tree))
}
