package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l, r := 0, m*n-1

	for l <= r {
		mid := (l + r) / 2
		midRow := mid / n
		midCol := mid % n
		if matrix[midRow][midCol] == target {
			return true
		} else if matrix[midRow][midCol] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 4
	fmt.Println(searchMatrix(matrix, target))
}
