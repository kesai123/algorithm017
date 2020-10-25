package main

import (
	"fmt"
	"sort"
)

//贪心

func main(){
	g := []int{1,2,3}
	s := []int{1,1}
	ret1 := findContentChildren(g,s)
	fmt.Println("ret1 = ", ret1)

	g = []int{1,2}
	s = []int{1,2,3}
	ret2 := findContentChildren(g,s)
	fmt.Println("ret2 = ", ret2)
}


func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0

	for i,j := 0,0; i<len(g) && j<len(s); {
		if g[i] <= s[j]{
			count++
			i++
			j++
		}else {
			j++
		}
	}
	return count
}
