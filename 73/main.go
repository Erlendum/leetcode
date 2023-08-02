package main

import "fmt"

func setZeroes(matrix [][]int) {
	mi := make(map[int]bool, 0)
	mj := make(map[int]bool, 0)

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				mi[i] = true
				mj[j] = true
			}
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			if mi[i] || mj[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	fmt.Println(matrix)
	setZeroes(matrix)
	fmt.Println(matrix)
}
