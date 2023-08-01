package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)
	tmp := make([][]int, n)
	for i := range tmp {
		tmp[i] = make([]int, n)
	}

	for i := range tmp {
		copy(tmp[i], matrix[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[j][n-i-1] = tmp[i][j]
		}
	}

}

func main() {
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(a)
	rotate(a)
	fmt.Println(a)
}
