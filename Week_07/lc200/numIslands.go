package main

func numIslands(grid [][]byte) int {
	if len(grid)==0 {return 0}
	//sol := newSolutionBFS(grid)
	sol := newSolutionDFS(grid)
	return sol.numIslands()
}

type SolutionBFS struct{
	grid [][]byte
	q [][2]int
	nr int
	nc int
}

func newSolutionBFS(g [][]byte) *SolutionBFS {
	sol := SolutionBFS{
		g,
		make([][2]int,0),
		len(g),
		len(g[0]),
	}
	return &sol
}

func (sol *SolutionBFS) numIslands() int{
	num := 0
	for i := range sol.grid {
		for j := range sol.grid[i] {
			if sol.grid[i][j] != '1' {continue}
			num++
			sol.bfs(i, j)
		}
	}
	return num
}

func (sol *SolutionBFS) bfs(i, j int){
	sol.q = append(sol.q, [2]int{i,j})
	for len(sol.q)>0 {
		r, c := sol.q[0][0], sol.q[0][1]
		sol.q = sol.q[1:]
		neighbors := [][2]int{{r-1,c}, {r+1,c}, {r,c-1}, {r,c+1}}
		for k := range neighbors {
			sol.subProcess(neighbors[k][0], neighbors[k][1])
		}
	}
}

func (sol *SolutionBFS) subProcess(r, c int){
	if r >= 0 && r <= sol.nr-1 && c >=0 && c <= sol.nc-1 {
		if sol.grid[r][c] == '1' {
			sol.q = append(sol.q, [2]int{r,c})
			sol.grid[r][c] = '0'
		}
	}
}

type SolutionDFS struct{
	grid [][]byte
	nr int
	nc int
}

func newSolutionDFS(g [][]byte) *SolutionDFS {
	sol := SolutionDFS{
		g,
		len(g),
		len(g[0]),
	}
	return &sol
}

func (sol *SolutionDFS) numIslands() int{
	num := 0
	for i := range sol.grid {
		for j := range sol.grid[i] {
			if sol.grid[i][j] == '1' {
				sol.dfs(i, j)
				num++
			}
		}
	}
	return num
}

func (sol *SolutionDFS) dfs(i, j int){
	if i<0 || i>sol.nr-1 || j<0 || j>sol.nc-1 || sol.grid[i][j]=='0' {return}
	sol.grid[i][j] = '0'
	neighbors := [][2]int{{i-1,j}, {i+1,j}, {i,j-1}, {i,j+1}}
	for k := range neighbors {
		sol.dfs(neighbors[k][0], neighbors[k][1])
	}
}