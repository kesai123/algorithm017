package main

import "fmt"

func main(){
	intervals1 := [][]int {{1,3},{2,6},{8,10},{15,18}}
	ret1 := merge(intervals1)
	fmt.Println(ret1)
	intervals2 := [][]int {{1,4},{4,5}}
	ret2 := merge(intervals2)
	fmt.Println(ret2)
}
