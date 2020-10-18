package main

import "fmt"

func main() {
	ret := combine(4,3)
	for _, ret1 := range ret {
		for _, ret2 := range ret1{
			fmt.Printf("%d ", ret2)
		}
		fmt.Printf("\n")
	}
}

func combine(n int, k int) [][]int {
	sol := NewSolution(n,k)
	if k <= 0 || n < k {
		return sol.ret
	}
	sol.Dfs(1)
	return sol.GetResult()
}

type Solution struct {
	n int
	k int
	ret [][]int
	path []int
}

func NewSolution(n, k int) *Solution{
	sol := new (Solution)
	sol.k = k
	sol.n = n
	sol.ret = [][]int{}
	sol.path = []int{}
	return sol
}

func (sol *Solution)GetResult() [][]int {
	return sol.ret
}

func (sol *Solution)Dfs(start int) {
	//剪枝
	if(len(sol.path)+(sol.n-start+1) < sol.k){
		return
	}
	//递归终止
	if len(sol.path) == sol.k {
		comb := make([]int, sol.k)
		copy(comb, sol.path)
		sol.ret = append(sol.ret, comb)
		return
	}
	//选择当前值 需要reverse path
	sol.path = append(sol.path, start)
	sol.Dfs(start+1)
	sol.path = sol.path[:len(sol.path)-1]
	//不选择当前值
	sol.Dfs(start+1)
}
