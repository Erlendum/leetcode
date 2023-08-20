package main

import "fmt"

func searchLand(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'

	searchLand(grid, i+1, j)
	searchLand(grid, i-1, j)
	searchLand(grid, i, j+1)
	searchLand(grid, i, j-1)
}

func numIslands(grid [][]byte) int {
	var res int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				res++
				searchLand(grid, i, j)
			}
		}
	}
	return res
}

func main() {
	grid := [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}}
	fmt.Println(numIslands(grid))
}
