package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	visited := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		visited[i] = make([]bool, len(matrix[i]))
	}
	res := make([]int, 0)
	i, j := 0, 0
	visited[i][j] = true
	res = append(res, matrix[i][j])
	prio := -1
	for {
		if j+1 < len(matrix[i]) && !visited[i][j+1] && (prio == 0 || prio == -1) {
			visited[i][j+1] = true
			res = append(res, matrix[i][j+1])
			j++
			prio = 0
		} else if i+1 < len(matrix) && !visited[i+1][j] && (prio == 1 || prio == -1) {
			visited[i+1][j] = true
			res = append(res, matrix[i+1][j])
			i++
			prio = 1
		} else if j-1 >= 0 && !visited[i][j-1] && (prio == 2 || prio == -1) {
			visited[i][j-1] = true
			res = append(res, matrix[i][j-1])
			j--
			prio = 2
		} else if i-1 >= 0 && !visited[i-1][j] && (prio == 3 || prio == -1) {
			visited[i-1][j] = true
			res = append(res, matrix[i-1][j])
			i--
			prio = 3
		} else if prio != -1 {
			prio = -1
		} else {
			break
		}
	}
	return res
}

func main() {
	a := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	fmt.Println(spiralOrder(a))
}
