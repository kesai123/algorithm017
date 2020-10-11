package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var input = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	output := groupAnagrams(input)

	for i := range(output) {
		s := "["
		for j:=range output[i] {
			s = s + output[i][j] + " "
		}
		s += "]\n"
		fmt.Print(s)
	}
}

func groupAnagrams(strs []string) [][]string {
	ret := [][]string{}
	m := map[string]int{}

	for _, str := range strs {
		key := myHash1(str)
		idx, ok := m[key]
		if ok {
			ret[idx] = append(ret[idx], str)
		} else {
			m[key] = len(ret)
			ret = append(ret, []string{str})
		}
	}
	return ret
}

// 基于字符排序构建key
func myHash1(s string) string {
	v := []byte(s)
	sort.Slice(v, func(i,j int)bool{return v[i]<v[j]})
	return string(v)
}

// 基于字符count map构建key
func myHash2(s string) string {
	var ret string
	v := [26]int{0}
	for _, b := range s{
		v[b-'a']++
	}
	for i := range v{
		ret += strconv.Itoa(v[i])
	}
	return ret
}