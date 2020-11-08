package main

func minPathSum(grid [][]int) int {
	for i := range grid {
		for j := range grid[i] {
			if i==0 && j==0 {
				continue
			} else if i==0 {
				grid[i][j] += grid[i][j-1]
			} else if j==0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i][j-1], grid[i-1][j])
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a <= b {return a} else {return b}
}