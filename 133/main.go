package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func dfsClone(node *Node, m map[*Node]*Node) *Node {
	neighbours := make([]*Node, 0)
	nodeCloned := new(Node)
	nodeCloned.Val = node.Val

	m[node] = nodeCloned

	for _, v := range node.Neighbors {
		if vCloned, ok := m[v]; ok {
			neighbours = append(neighbours, vCloned)
		} else {
			neighbours = append(neighbours, dfsClone(v, m))
		}
	}
	nodeCloned.Neighbors = neighbours
	return nodeCloned
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	if len(node.Neighbors) == 0 {
		nodeCloned := new(Node)
		nodeCloned.Val = node.Val
		return nodeCloned
	}
	m := make(map[*Node]*Node)

	return dfsClone(node, m)
}
