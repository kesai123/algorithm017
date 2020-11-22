package main

import (
	"sort"
)

type sortable [][]int
func (a sortable)Len() int {return len(a)}
func (a sortable)Swap(i, j int) {a[i],a[j] = a[j],a[i]}
func (a sortable)Less(i, j int) bool {return a[i][0]< a[j][0]}

func merge(intervals [][]int) [][]int {
	ret := make([][]int, 0)
	if len(intervals) == 0 {return ret}
	sort.Sort(sortable(intervals))
	for i := range intervals {
		l, r := intervals[i][0], intervals[i][1]
		if i == 0 || ret[len(ret)-1][1] < l {
			ret = append(ret, intervals[i])
		} else {
			ret[len(ret)-1][1] = maxInt(ret[len(ret)-1][1], r)
		}
	}
	return ret
}

func maxInt(a,b int) int{
	if a >= b {return a} else {return b}
}